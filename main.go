package main


import (

    "github.com/gin-gonic/gin"
    "RealTLocation/controllers/orders"
    "RealTLocation/location_parser/loggers"
    "RealTLocation/grpc/client"
    "github.com/gorilla/cors"
    "RealTLocation/models/models"
    "fmt"
    "os"
)

var (
    applicationHost = os.Getenv("APPLICATION_HOST")
    applicationPort = os.Getenv("APPLICATION_PORT")
)

func StartGRPCServer(){
    loggers.DebugLogger.Println("Starts gRPC Server....")
    client.StartGRPCServer()
}

func main(){
    models.database.AutoMigrate(&models.Order{}, &models.Goods{} &models.destinationAddress{})
    router := gin.Default()
    allowedOrigins := []string{fmt.Sprintf("%s:%s", applicationHost, applicationPort)}
    allowedMethods := []string{"PUT", "POST", "GET", "PUT", "DELETE"}
    allowedHeaders := []string{"*"}
    cors.CORS(allowedOrigins, allowedMethods, allowedHeaders, router)
    router.GET("/delivered/orders/", orders.GetDeliveredOrders)
    go StartGRPCServer()
    router.Run(fmt.Sprintf(":%s", applicationPort))
}

