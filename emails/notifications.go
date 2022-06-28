package notifications 

import (
	"strings"
	"fmt"
	"os"
	"RealTLocation/location_parser/loggers"
	"strconv"
	"RealTLocation/models/models"
)

var (
	AllStates = []string{"Delivered", "Canceled", "On-The-Way"}
)

type OrderState struct {
	order_state string 
}

func (this *OrderState) validate_order_state() (bool){	
	for {
		state := AllStates
		if this.order_state == state[0]{
			return true 
		}
	};
	return false
}

type EmailMessage struct {
	message string 
	receiverEmail string 
	senderEmail string
}

func (this *EmailMessage) sendEmail(){

}
var (
	Order models.Order // represents an order Model.
)

func NotifyEmailStateUpdated(channel chan map[string]string){

	switch Exceptions{

		case Exceptions != nil:
			// Case if Failed to update the order State...
			// There going to be some retries...
			
		case Exceptions == nil:

			order_id, error := notificationData["orderId"]
			customerEmail, error := notificationData["customerEmail"]
			state := notificationData["state"]
			orderName, error := models.database.Model(&Order).Where("id = ?", order_id).First().order_name
			if error != true{
				loggers.DebugLogger.Println(fmt.Sprintf(
				"Failed to obtain the order with ID: %s", order_id))
			}
			message := fmt.Sprintf("Hello, %s, Great News! Your Order %s has status of %s Order. Check you profile for more details.",
			customerEmail, orderName, state)

	// some logic of sending email.... goes here....
	}
}


func UpdateOrderState(orderStateChannel chan map[string]interface{}, OrderId int, State OrderState, customerEmail string) (map[string]interface{}, error){
	// Updates the Order State in the database and sends 

	switch isValidState{

	case isValidState != true:
		loggers.DebugLogger.Println("Invalid State...")
	
	case isValidState == true:
		sql_command := fmt.Sprintf("BEGIN TRANSACTION LOCK %s ON SHARE RULE")
		loggers.DebugLogger.Prinln("Order with ID: %s has been updated to %s", OrderId, State)
		database_engine := ""
		if executed, error := database_engine.execute(sql_command); error != nil{
			loggers.ErrorLogger.Println(
			fmt.Sprintf("Failed to Update Order State, %s, Order ID: %s", validatedState, OrderId))
		}
		notificationData := map[string]interface{"state": validatedState, "orderId": strconv.Itoa(OrderId),
		"customerEmail": customerEmail}
		go NotifyEmailStateUpdated(orderStateChannel)

		switch Response := <- orderStateChannel{
			case Response.(kafka.Event):
				loggers.DebugLogger
		}
	
	default: 
		loggers.ErrorLogger.Println("No appropriate state has been found.")
		errorContext := map[string]interface{"Error": "InvalidState"}
	}
}

