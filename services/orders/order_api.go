package order_api 

import (
	_ "strings"
	"fmt"
	_ "os"
	"RealTLocation/location_parser/loggers"
	"strconv"
	"RealTLocation/models/models"
	"errors"
	"RealTLocation/emails/notifications"
)

var (
	AllStates = []string{"Delivered", "Canceled", "On-The-Way"}
)

func UpdateOrderState(OrderId int, State string, customerEmail string) (bool, error){
	// Updates the Order State in the database and sends 

	// If Order state has been passed valid... Executing Update State Transaction with LOCK in order to safe the order state.
	var order models.Order
	models.database.Model(
	&order).Where("id = ?", OrderId).First().Raw("FOR UPDATE")// 1. Receives Order,
	   // 2. Lock It in order to restrict any updates from over services, 
	   // 3. Updates object and Commit Transaction.

	loggers.DebugLogger.Prinln("Order with ID: %s has been updated to %s", OrderId, State)
	notificationData := map[string]string{"state": State, "orderId": strconv.Itoa(OrderId),
	"customerEmail": customerEmail}
	sended, error := notifications.NotifyEmailOrderStateUpdated(notificationData)

	switch sended{

		case error != nil:
			loggers.ErrorLogger.Println(fmt.Sprintf("Failed to send Email Notification. Email: %s, State: %s",
			customerEmail, State))
			return false, nil 

		case sended != true && error == nil:
			loggers.DebugLogger.Println(fmt.Printf("Notification has been sended successfully."))
			return true, nil
	}
	return false, errors.New(
	"Unknown State, did not receive proper state of the operation.")
	}



