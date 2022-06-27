package main 

import (
	"errors"
	"os"
	"net/http"
	"fmt"
	"logging"
	"RealTLocation/location_parser/exceptions"
	"json"
	"sync"
	"reflect"
	"github.com/go-gocron/gocron"
)

var (
	ErrorLogger *logging.GetLogger 
	DebugLogger *logging.GetLogger 
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
}

func (this *ParsedLocation) convertToJson() (string, error){
}



func parse_customer_location(deviceInfo *DeviceInfo) (map[string]string, error){
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





