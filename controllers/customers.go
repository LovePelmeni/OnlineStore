package customers

import (
	"models"
	"sync"
	"time"
)

var customer models.Customer
var customers []models.Customer

type CustomerData struct {
	sync.Mutex
	Username string 
	Email string 
	Password string 
	createdAt time.Time 
}

func createRemotePaymentController(){
}

func CreateCustomerController(customer_data *CustomerData){

}

func UpdateCustomerController(){

}

func DeleteCustomerController(){

}