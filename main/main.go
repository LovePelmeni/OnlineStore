package main

import (
  "OrderCheckout/orders/orders"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
  "fmt"
  "os"
)

var (
    // Service Allowed Hosts Applications...
    port = os.Getenv("APPLICATION_PORT")

)

func main(){

  router := gin.Default()
  allowedOrigins := []string{}
  allowedMethods := []string{}
  allowedHeaders := []string{"*"}
  allowCredentials := true

  config := cors.Config{}
  config.AllowOrigins = allowedOrigins 
  config.AllowHeaders = allowedHeaders 
  config.AllowMethods = allowedMethods 
  config.AllowCredentials = allowCredentials 

  handleControllers := router.Group("/order/")
  {
  handleControllers.DELETE("deny/:orderId/", checkout.RejectConfirmIncomingOrder)
  handleControllers.POST("accept/:orderId/", checkout.AcceptConfirmIncomingOrder)
  }
   

  getterControllers := router.Group("/orders/")
  {
    getterControllers.GET("list/:customerId/", checkout.GetAllOrders)
    getterControllers.GET("retrieve/:orderId/:customerId/", checkout.GetOrder)
  }
  router.Run(fmt.Sprintf(":%", port))
}
