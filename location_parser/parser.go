package main

import (
	"RealTLocation/location_parser/exceptions"
	"RealTLocation/models/models"
	"errors"
	"fmt"
	"json"
	"logging"
	"models"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"sync"
	"json"
	"xml"
	"github.com/go-gocron/gocron"
)

var (
	ErrorLogger *logging.GetLogger 
	DebugLogger *logging.GetLogger 
	AllStates = []string{"Delivered", "Preparing", "On-the-Way"}
)


type IpAddress struct {
	mutex sync.Mutex 
	value string 
}

func (this *IpAddress) validateIpAddress() (string, error){
	return "", nil 
}

func (this *IpAddress) getValue(){
}


type DeviceInfo struct {
	mutex sync.Mutex 
	device_name string 
	device_ip_address IpAddress
}

func (this *DeviceInfo) validate() (*DeviceInfo, error){
	reflect.ValueOf(&this).Elem()
	return this, nil 
}


type ParsedLocation struct {
	mutex sync.Mutex
	latitude string 
	longitude string 
}


func (this *ParsedLocation) convertToXml() (string, error){
	return xml.Marshal(this)
}

func (this *ParsedLocation) convertToJson() (string, error){
	return json.Marshal(this)
}


func parse_customer_location(deviceInfo *DeviceInfo) (map[string]interface{}, error){
	return 
}


func send_kafka_location_event(location string) (bool, interface{}){
	return bool(true), nil
}

func ParseLocationTask(deviceInfo *DeviceInfo) (interface{}){

	raw_parsed_location, location_error := parse_customer_location(deviceInfo)

	switch raw_parsed_location{

		case location_error.(*exceptions.CustomerIsOffline):
			DebugLogger.Println(fmt.Sprintf(
			"Customer Is Offline. Removing task. %s", location_error.Error()))
			gocron.Remove(ParseLocationTask)
		
		default:
		
			ParsedLocation.mutex.Lock()
			location := ParsedLocation{
				longitude: raw_parsed_location["longitude"],
				latitude: raw_parsed_location["latitude"],
			}
			converted_data, xml_error := location.convertToXml()
			if xml_error != nil{
				ErrorLogger.Println(
				fmt.Sprintf("XML CONVERTION ERROR: %s", xml_error))
				
			}
			defer ParsedLocation.mutex.UnLock()
			send_kafka_location_event(converted_data)

		case location_error.(*exceptions.ParseError):
			ErrorLogger.Prinln(fmt.Sprintf("Parser Error. %s", location_error.Error()))
			// locking task... and sending report issue...
	}
}

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
		ErrorLogger.Println(
		fmt.Sprintf("Failed to Update Order State, %s, Order ID: %s", validatedState, OrderId))
	}
	notificationData := map[string]interface{"state": validatedState, "orderId": orderId,
    "customerEmail": customerEmail}
	orderStateChannel <- notificationData
}


func start_consuming(deviceInfo *DeviceInfo) (bool, interface{}){

	validatedDeviceInfo, validateException := deviceInfo.validate()
	switch validateException {
		case validateException.(exceptions.InvalidDeviceInfo):
			ErrorLogger.Println(fmt.Sprintf("Invalid Device Info %s", json.Marshal(*deviceInfo)))
			return false, validateException
		}

	error := ParseLocationTask(validatedDeviceInfo)
	if error != nil{
		ErrorLogger.Println(fmt.Sprintln(
		"Failed To Initialize task... %s", errors.New("Failed to start Task.")))
		return false, error
	}
}




