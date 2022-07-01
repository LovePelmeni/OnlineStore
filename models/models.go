package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	_ "strconv"
	"time"
	"fmt"

  "gorm.io/datatypes"

)

var (

dbHost = os.Getenv("DATABASE_HOST")
dbUser = os.Getenv("DATABASE_USER")
dbPassword = os.Getenv("DATABASE_PASSWORD")
dbPort = os.Getenv("DATABASE_PORT")
dbName = os.Getenv("DATABASE_NAME")



databaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
database, error = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
	DisableForeignKeyConstraintWhenMigrating: true,
}))

var (
  OrderStates = []string{"Pre-Confirmed", "Confirmed", "Canceled", "Delivered"}
)
var order Order 

type DestinationAddress struct {
  // structure reprensents the destination Address

	id int64
	latitude string
	longitude string

  city string
  country string

	street string
	house string

}


type Order struct {
  gorm.Model
	id int64
	order_name string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	goods datatypes.JSON
	createdAt time.Time `gorm:"DATE DEFAULT CURRENT_DATE"`
	destinationAddress datatypes.JSON
	customerEmail string `gorm:"VARCHAR(10) NOT NULL`
	state string `gorm:"VARCHAR(10) NOT NULL"`
}



type OrderCustomerCredentials struct {
    // Structure represents the
    customerEmail string
    phoneNumber string
}



// Creates an order.

func createOrder(OrderData map[string]string,

 
  customerCredentials OrderCustomerCredentials,
  destinationAddress DestinationAddress) (*Order){
  orderTime := time.Now()

  orders := database.Model(&Order{}).Where("HAVING id").Statement.ReflectValue
  order_name := fmt.Sprintf("Order-%s", orders)
  serializedCustomerCredentials := datatypes.JSON(
  []byte(fmt.Sprintf(`{"email": "%s", "phone": "%s"}`,
  customerCredentials.customerEmail, customerCredentials.phoneNumber)))
  serializedDestinationAddress := datatypes.JSON(
    []byte(fmt.Sprintf(`{"state": "%s", "city": "%s",
    "street": "%s", "house": "%s"}`,  )),
  )

  newOrder := Order{
  customerCredentials: serializedCustomerCredentials,
  destinationAddress: serializedDestinationAddress,
  order_name: order_name,
  createdAt: orderTime,
  state: fmt.Sprintf("%s", OrderStates[0]), // current state of the order. default "pre-confirmed"
}


  database.Save(&newOrder)
  return &newOrder

}



// Deletes order
func deleteOrder(orderId string) (bool){
  database.Model(&Order{}).Where("id = ?", orderId).Delete(order)
  return true 
}



// Updated Order.
func UpdateOrder(NewOrderData Order, orderId string) (bool){

  // UnEditableFields := []string{"order_name", "createdAt"}
  // UpdateSequence := []string{}
  // Element := reflect.TypeOf(order)
  // for value := 1; value < Element.NumField(); value ++{
  //       Field := Element.Field(value)
  //       reflect.Append(reflect.ValueOf(UpdateSequence), 
  //       reflect.ValueOf(Field.Name), reflect.ValueOf(Field)) // appends elements in order : `key`, `value` as requuires to properly update object
  // } 
  // orderObject := database.Model(order).Where("id = ?", 
  // orderId).First(order).Update(UpdateSequence)
  // return true, nil 
  return true
}





type Goods struct {
  // structure represents the info about the delivery product
	id int64
	Names *[]string  // names of the products that has been purchased..
  Product_Ids *[]int // identifiers of the products that has been purchased...
	Quantity string
	TotalPrice float64
}
