package notifications 

import (
	"strings"
	"fmt"
	"os"
	"RealTLocation/location_parser/loggers"
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
	notificationData := <- channel
	order_id, error := notificationData["orderId"]
	customerEmail, error := notificationData["customerEmail"]
	state := notificationData["state"]
	orderName, error := models.database.Model(&Order).Where("id = ?", order_id).First().order_name
	if error != true{
		DebugLogger.Println(fmt.Sprintf(
		"Failed to obtain the order with ID: %s", order_id))
	}
	message := fmt.Printf("Hello, %s, Great News! Your Order %s has status of %s Order. Check you profile for more details.",
	customerEmail, orderName, state)
	// some logic of sending email.... goes here....
}

func UpgradeOrderState(orderStateChannel chan map[string]interface{}, transactionData map[string]interface{}, OrderId int, State OrderState, customerEmail string){
	// Updates the Order State in the database and sends 
	validatedState := State.validate_order_state()
	sql_command := fmt.Sprintf("BEGIN TRANSACTION LOCK %s ON SHARE RULE")
	DebugLogger.Prinln("Order with ID: %s has been updated to %s", OrderId, State)
	database_engine := ""
	if executed, error := database_engine.execute(sql_command); error != nil{
		loggers.ErrorLogger.Println(
		fmt.Sprintf("Failed to Update Order State, %s, Order ID: %s", validatedState, OrderId))
	}
	notificationData := map[string]interface{"state": validatedState, "orderId": orderId,
    "customerEmail": customerEmail}
	orderStateChannel <- notificationData
}