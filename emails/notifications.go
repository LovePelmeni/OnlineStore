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

func sendedEmailWebhook(){
	// waits for the email webhook
}

func sendEmailNotification(channel chan map[string]bool, message string, customerEmail string, senderEmail string) (bool, error){
	// some logic of sending email...
	go sendedEmailWebhook(channel)
}

func NotifyEmailOrderStateUpdated(notificationData map[string]string) (bool, error){
	deliveryEmailChannel := make(chan map[string]bool)
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

	go sendEmailNotification(message, customerEmail, senderEmail)
	deliveryResponse := <- deliveryEmailChannel

	switch deliveryResponse {
		case deliveryResponse["delivered"]:
			loggers.ErrorLogger.Println("Failed to deliver Email Notification...")
			// some continue.
		
		case deliveryResponse["delivered"]:
			loggers.DebugLogger.Println("Message has been delivered successfully.")
}

