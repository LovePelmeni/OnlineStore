package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

var (
    // Service Allowed Hosts Applications...
    port := os.Getenv("APPLICATION_PORT")

)

func main(){

  router := gin.Default()
  allowedOrigins := []string{}
  allowedMethods := []string{}
  allowedHeaders := []string{"*"}
  allowCredentials := true

  router.Group("/order/"){
    router.DELETE("deny/:orderId/", checkout.RejectConfirmIncomingOrder),
    router.POST("accept/:orderId/", checkout.AcceptConfirmIncomingOrder),
  }
  router.Group("/orders/"){
    router.GET("list/:customerId/", checkout.GetAllOrders),
    router.GET("retrieve/:orderId/:customerId/", checkout.GetOrder),
  }
  router.Run(fmt.Sprinf(":%", port))
}
