package parser 

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
	"net"
	"RealTLocation/location_parser/kafka"
	"url"
	"ioutil"
	"RealTLocation/location_parser/loggers"
)

var (
	ErrorLogger *logging.GetLogger 
	DebugLogger *logging.GetLogger 
	AllStates = []string{"Delivered", "Preparing", "On-the-Way"}
	geoPositionAPIKey = os.Getenv("GEO_POSITIION_API_KEY")
)


type MacAddress struct {
	mutex sync.Mutex 
	value string 
}

func (this *MacAddress) validateMacAddress() (bool, error){
	if valid := net.ParseIP(this.value); valid != nil{
		return true, nil 
	}
	return false, errors.New("Invalid IP Address.")
}


type DeviceInfo struct {
	mutex sync.Mutex 
	device_name string 
	macAddress MacAddress 
}

func (this *DeviceInfo) validate() (*DeviceInfo, error){
	if validated, error := this.macAddress.validateMacAddress(); validated != true || error != nil{
		loggers.DebugLogger.Println(fmt.Sprintf("Invalid Mac Address Occurred.. %s", error.Error()))
	}
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

	var parsedLocationLink ParsedLocation
	urlAddress := url.Parse(fmt.Sprintf("https://pos.ls.hereapi.com/positioning/v1/locate"))
	macAddress := deviceInfo.macAddress.value

	position := json.Marshal(map[string]interface{}{
		"wlan": []struct{mac string}{
			{mac: macAddress},
		},
	})
	queryParams := urlAddress.Query()
	queryParams.Set("apiKey", geoPositionAPIKey)

	response, http_request_error := http.Post(urlAddress.String(), "application/json", position)
	if http_request_error != nil {
		loggers.ErrorLogger.Println("Failed to Send Http Request..")
		return nil, exceptions.ParserError 
	}

	if readedData, error := ioutil.ReadAll(response.Body); readedData != nil && error != nil{
		decodedData := json.Unmarshal(readedData, &parsedLocationLink) // decoding data in according to the structure....
		loggers.DebugLogger.Println("Location has been parsed successfully.. Longitude: %s, Latitude: %s",
		decodedData["longitude"], decodedData["latitude"])
		return readedData, nil 

	} else {
		loggers.ErrorLogger.Println("Failed to parse Location....")
		return nil, exceptions.ParserError
	}
}


func send_kafka_location_event(topic string, partition string, location string) (bool, interface{}){

	kafka_client := kafka.KafkaClient.createProducer()
	KafkaResponseEventChannel := make(chan kafka.Event) // event for listen for kafka event response...
	go kafka_client.ProduceEvent(topic, partition, location, KafkaResponseEventChannel)

	sended := <- KafkaResponseEventChannel
	switch sended {
		case sended.(kafka.Event):
			loggers.DebugLogger.Println(fmt.Sprintf("Event has been delivered successfully."))
			return true, nil 

		case sended.(kafka.Event) == nil:
			loggers.ErrorLogger.Println(
				fmt.Sprintf("Event has been missed.. Responded with None"))
				return false, nil
		default: 
			loggers.DebugLogger.Println(
			"No Available Cases For Kafka Response has been detected. Failing..")
			return false, nil
		}
	}

func ParseLocationTask(deviceInfo *DeviceInfo, parsedLocationChannel chan map[string]interface{}){

	parsed_location, location_error := parse_customer_location(deviceInfo)

	switch parsed_location{

		case location_error.(*exceptions.CustomerIsOffline):
			DebugLogger.Println(fmt.Sprintf(
			"Customer Is Offline. Removing task. %s", location_error.Error()))
			gocron.Remove(ParseLocationTask) /// removing periodic task...
		
		default:
			parsedLocationChannel <- parsed_location

		case location_error.(*exceptions.ParseError):
			ErrorLogger.Prinln(fmt.Sprintf("Parser Error. %s", location_error.Error()))
			// locking task... and sending report issue...
	}
}

func startParseLocationTask(deviceInfo *DeviceInfo) (bool, interface{}){

	validatedDeviceInfo, validateException := deviceInfo.validate()
	switch validateException {
		case validateException.(exceptions.InvalidDeviceInfo):
			ErrorLogger.Println(fmt.Sprintf("Invalid Device Info %s", json.Marshal(*deviceInfo)))
			return false, validateException
		}
	channel := make(chan map[string]interface{})
	go gocron.Every(5).Second().Do(ParseLocationTask, validatedDeviceInfo, channel)
	return true, nil
}




