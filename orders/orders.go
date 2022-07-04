package orders

import (
  "github.com/LovePelmeni/OnlineStore/OrderCheckout/kafka"
  "github.com/LovePelmeni/OnlineStore/OrderCheckout/models"
  emails "github.com/LovePelmeni/OnlineStore/OrderCheckout/emails"
  "github.com/gin-gonic/gin"
  _ "errors"
  "net/http"
  "errors"
  "fmt"
  "log"
)

var order models.Order


// Order API Controllers goes there...


// loggers 
var (
  DebugLogger *log.Logger 
  ErrorLogger *log.Logger 
  WarningLogger *log.Logger 
  InfoLogger *log.Logger 
)

// Kafka Producers.
var (
  kafkaBackend = kafka.KafkaBackend{}
)

// var ProducerEntity, error_ = kafkaBackend.CreateKafkaProducer()

// http endpoint for accepting an orders.
func AcceptIncomingOrderController(context *gin.Context){
    AcceptIncomingOrder(context.Param("OrderId"))
    context.JSON(http.StatusAccepted, gin.H{"status": "accepted"}) // status code responsible to notify that order has been confirmed..
}

// http endpoint for rejecting an orders.
func RejectIncomingOrderController(context *gin.Context){
  orderId := context.Param("orderId")
  RejectIncomingOrder(orderId)
  context.JSON(http.StatusNotAcceptable, gin.H{"status": "rejected"}) // status code responsible to notify that it has been failed to proceed.
}

// http endpoint for getting all customer Orders.
func GetAllCustomerOrdersController(context *gin.Context){

  customerId := context.Param("customerId")
  query := models.Database.Model(order).Where(
  fmt.Sprintf("customer.id = %s", customerId))
  context.JSON(http.StatusOK, gin.H{"orders": query})
}

// http endpoint for getting specific customer Order.
func GetOrderController(context *gin.Context){

  customerId := context.Param("customerId")
  orderId := context.Param("orderId")
  queryObj := models.Database.Model(order).Where(
  "customer.id = ? AND id = ?", customerId, orderId)
  context.JSON(http.StatusOK, gin.H{"order": queryObj})
}



type OrderData struct {
  CustomerEmail string 
}
var orderData OrderData 

// API Handlers goes there...



// Method For Processing Rejection of the Order.

func RejectIncomingOrder(orderId string) (bool, error){
  // Deletes the order from the database and sends event back to the main service.

  orderObject := models.Database.Table("order").Where(
  "id = ?", orderId).Scan(&orderData)
  if orderObject.Error != nil {return false, errors.New("Order Object Does Not Exist.")}

  customerEmail, err := orderObject.Get("CustomerEmail")
  if err != false {return false, errors.New("Customer Email is Empty.")}
  // sends notification Email to let client know that the order has been rejected..
  sended, error := emails.NotifyEmailOrderRejected(fmt.Sprintf("%v", customerEmail), orderId)

  // // message for the kafka event.
  // message := fmt.Sprintf("Hello, %s, Unfortunately your order ID: %s has been rejected." +
  // "Go Check Your Profile for More Info.", customerEmail, orderId)

  if sended != true || error != nil {return false, errors.New("Failed to Send Email Notification.")}
  orderObject.Delete(order)

  // sending kafka event to `rejectedOrder` topic.
  // Producer := kafka.KafkaProducer{Producer: ProducerEntity, Topic: "rejected"} // creating a producer with following topic to let it send events to that topic.
  // responded, error := Producer.SendRejectOrderEvent(message, orderId) // sending event.
  // if responded == true && error == nil {return true, nil}else{return false, errors.New("Failed to Send Rejection Event.")}
  return true, nil
}



// Method For Processing Acceptance of the Order.

func AcceptIncomingOrder(orderId string) (bool, error){
    // Updates state of the order to confirmed and sends kafka event back to the main service.
    orderObject := models.Database.Model(&order).Where(
      "id = ?", orderId).First(order)
    customerEmail, error := orderObject.Get("customerEmail")
    if error != true {return false, errors.New("Empty Customer Email.")}
    emails.NotifyEmailOrderAccepted(fmt.Sprintf("%v", customerEmail), orderId)
    DebugLogger.Println("Transaction on the way confirm.")
    return true, nil 
}




