package orders

import (
  kafka_api "OrderCheckout/kafka/kafka"
  "OrderCheckout/models/models"
  "OrderCheckout/checkout/checkout"
  "OrderCheckout/loggers/loggers"
  "OrderCheckout/emails/emails"
  "github.com/gin-gonic/gin"
  "errors"
  "net/http"
  "gorm.io/datatypes"
  "encoding/json"
)

var order models.IncomingOrder


func AcceptIncomingOrderController(context *gin.Context){
    // http endpoint for accepting an orders.
    AcceptConfirmIncomingOrder(context.Param("OrderId"))
    context.JSON(http.StatusAccepted, gin.H{}) // status code responsible to notify that order has been confirmed..
}

func RejectIncomingOrderController(context *gin.Context){
  // sends kafka event
  context.JSON(http.StatusNotAcceptable, gin.H{}) // status code responsible to notify that it has been failed to proceed.
}


func RejectConfirmIncomingOrder(orderId string) (bool, error){
  // Deletes the order from the database and sends event back to the main service.
  orderObject := models.database.Model(&order).Where(
  "id = ?", orderId).First()

  eventData := kafka_api.OrderEventData{
  message: "Order has been Rejected.", orderId: orderObject.id}
  customerEmail := json.Unmarshal([]byte(orderObject.customerCredentials)) // decoding JSON Field from the model
  
  sended, error := emails.NotifyEmailOrderRejected(customerEmail, orderObject.id)
  orderObject.Delete(order)
  kafka_api.SendRejectOrderEvent(eventData)
  return true, nil 
}

func AcceptConfirmIncomingOrder(orderId string) (bool, error){
    // Updates state of the order to confirmed and sends kafka event back to the main service.
    orderObject := models.database.Model(&order).Where(
      "id = ?", orderId).First()
    

    eventData := kafka_api.OrderEventData{
      message: "Order has been Accepted.",
      customerEmail: orderObject.customerEmail,
    }
    emails.NotifyEmailOrderAccepted(eventData)
    loggers.DebugLogger.Println("Transaction on the way confirm.")
    return true, nil 
}
