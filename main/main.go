package main

import (
  "github.com/LovePelmeni/OnlineStore/OrderCheckout/orders"
  "github.com/LovePelmeni/OnlineStore/OrderCheckout/kafka"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
  "fmt"
  "os"
)

var (
    // Service Allowed Hosts Applications...
    port = os.Getenv("APPLICATION_PORT")

)


func triggerKafkaConsumer(){
 // Triggers Kafka Consumer to start..

 backend :=  kafka.KafkaBackend{}
 consumer, error := backend.CreateKafkaConsumer() // Starts Kafka Consumer... to listen to Orders..
 if error != nil {panic("Failed To Start Consumer")}
 consumerPointer := kafka.KafkaConsumer{Consumer: consumer}
 go consumerPointer.ConsumeKafkaOrderEvents()
}


func main(){

  router := gin.Default()
  allowedOrigins := []string{"*"}
  allowedMethods := []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
  allowedHeaders := []string{"*"}
  allowCredentials := true

  config := cors.Config{}
  config.AllowOrigins = allowedOrigins 
  config.AllowHeaders = allowedHeaders 
  config.AllowMethods = allowedMethods 
  config.AllowCredentials = allowCredentials 


  go triggerKafkaConsumer() // starts Kafka Consumer..

  // urlpatterns for handling accept/deny operations for the orders.
  handleControllers := router.Group("/order/")
  {
  handleControllers.DELETE("deny/:orderId/", orders.RejectIncomingOrderController)
  handleControllers.POST("accept/:orderId/", orders.AcceptIncomingOrderController)
  }

  // getter urlpatterns, responsible for obtaining info about the orders.
  getterControllers := router.Group("/orders/")
  {
    getterControllers.GET("list/:customerId/", orders.GetAllCustomerOrdersController)
    getterControllers.GET("retrieve/:orderId/:customerId/", orders.GetOrderController)
  }

  router.Run(fmt.Sprintf(":%s", port))
}

