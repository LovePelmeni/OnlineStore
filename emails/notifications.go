package notifications 

import (
	"strings"
	"fmt"
	"os"
	"RealTLocation/location_parser/loggers"
	"strconv"
	"RealTLocation/models/models"
	"errors"
)

var (
	AllStates = []string{"Delivered", "Canceled", "On-The-Way"}
)


type EmailMessage struct {
	message string 
	receiverEmail string 
	senderEmail string
}

func (this *EmailMessage) sendEmail(){

}
var (
	Order models.Order // represents an order Model.
	OrderState models.OrderState

)

func sendEmailNotification(message string, customerEmail string, senderEmail string) (bool, error){
	// some logic of sending email...
}

func NotifyEmailOrderStateUpdated(notificationData map[string]string) (bool, error){

	senderEmail := os.Getenv("SENDER_EMAIL")
	order_id, error := notificationData["orderId"]

	customerEmail, error := notificationData["customerEmail"]
	state := notificationData["state"]

	orderName, error := models.database.Model(&Order).Where(
	"id = ?", order_id).First().order_name

	if error != true{
		loggers.DebugLogger.Println(fmt.Sprintf(
		"Failed to obtain the order with ID: %s", order_id))
	}

	message := fmt.Sprintf("Hello, %s, Great News! Your Order %s has status of %s Order. Check you profile for more details.",
	customerEmail, orderName, state)

	sendedStatus, error_ := sendEmailNotification(message, customerEmail, senderEmail)
	if error_ != nil {
		return false, error_
	}
	return sendedStatus, nil 
}

