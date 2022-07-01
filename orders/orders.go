package orders

import (
  models "OrderCheckout/models/models"
  checkout "OrderCheckout/checkout/checkout"
  loggers "OrderCheckout/loggers/loggers"
  "github.com/gin-gonic/gin"
  "errors"
)

var order models.IncomingOrder

func AcceptIncomingOrderController(context *gin.Context){
    // http endpoint for accepting an orders.
    AcceptConfirmIncomingOrder(context.Param("OrderId"))
    gin.JSON(http.StatusAccepted) // status code responsible to notify that order has been confirmed..
}

func RejectIncomingOrderController(gin *gin.Context){
  // sends kafka event
  gin.JSON(http.StatusNotAcceptable) // status code responsible to notify that it has been failed to proceed.
}


func RejectConfirmIncomingOrder() (bool, error){
  // Deletes the order from the database and sends event back to the main service.
  return NotifyRejected()
}

func AcceptConfirmIncomingOrder(orderId &order) (bool, error){
    // Updates state of the order to confirmed and sends kafka event back to the main service.
    loggers.DebugLogger.Println("Transaction on the way confirm.")

}
