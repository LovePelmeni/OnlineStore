package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	_ "strconv"
	"time"
	"fmt"
  "github.com/oleiade/reflections"
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


type RestrictedFields interface {
    // Interface allows to return fields that is not allowed to be updated.
    GetRestrictedFields() ([]string)
}


type DestinationAddress struct {
  // structure reprensents the destination Address
  gorm.Model
	id int64 
	latitude string `gorm:"VARCHAR(100) NOT NULL DEFAULT '0.0000'"`
	longitude string `gorm:"VARCHAR(100) NOT NULL DEFAULT '0.0000'"`

  city string `gorm:"VARCHAR(100) NOT NULL"`
  country string `gorm:"VARCHAR(100) NOT NULL"`

	street string `gorm:"VARCHAR(100) NOT NULL"`
	house int `gorm:"INTEGER NOT NULL"` 
}

func (this *Order) getRestrictedFields() ([]string){
  // returns restricted fields for Order Model..
}

func (this *Order) getAllFields() ([]string){
  // returns all model fields..
}


type Goods struct {
  // structure represents the info about the delivery product
  gorm.Model
	id int64
	Name string  `gorm:"VARCHAR(40) NOT NULL"`// names of the products that has been purchased..
  productId string `gorm:"VARCHAR(40) NOT NULL"` // identifiers of the products that has been purchased...
	Quantity string `gorm:"VARCHAR(50) NOT NULL DEFAULT '1'"`
	TotalPrice float64 `gorm:"NUMERIC(10, 2) NOT NULL"`
  currency string `gorm:"VARCHAR(10) NOT NULL DEFAULT 'usd'"`
}

func (this *Order) getRestrictedFields() ([]string){
  // returns restricted fields for Order Model..
}

func (this *Order) getAllFields() ([]string){
  // returns all model fields..
}


type Order struct {
  gorm.Model
	id int64
	order_name string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	goods Goods `gorm:"foreignKey Goods DEFAULT NULL"` // 
	createdAt time.Time `gorm:"DATE DEFAULT CURRENT_DATE"`
	destinationAddress DestinationAddress `gorm:"foreignKey destinationAddress NOT NULL"`
	customerEmail string `gorm:"VARCHAR(10) NOT NULL"`
	state string `gorm:"VARCHAR(10) NOT NULL"`
}

func (this *Order) getRestrictedFields() ([]string){
  // returns restricted fields for Order Model..
}

func (this *Order) getAllFields() ([]string){
  // returns all model fields..
}


type OrderCustomerCredentials struct {
    // Structure represents the
    customerEmail string
    phoneNumber string
}


func CreateRestrictColumnUpdateTrigger(
  // this Method Adds Restrictions for fields to Update.. 
  // Receives the model and list of fields which should not be updated.
model *gorm.DB, forbidden_columns []string) (bool){
  sql_first_part_command := `CREATE OR REPLACE FUNCTION restrictUpdateField RETURNS TRIGGER 
  LANGUAGE plpgsql AS
  $$BEGIN` 
  for {
    value := forbidden_columns 
    sql_first_part_command += fmt.Sprintf(`IF NEW.%s IS DISTINCT FROM OLD.%s` + 
    `NEW.%s = OLD.%s`, value, value)
  }
  sql_first_part_command += "END IF RETURN NEW; END;$$;"
  sql_second_part_command := fmt.Sprintf("CREATE TRIGGER update_restrictor BEFORE UPDATE %s" + 
  "FOR EACH ROW EXECUTE PROCEDURE restrictUpdateField", model.Name)
  error := model.Exec(
  sql_first_part_command + sql_second_part_command)
  if error.Statement.ReflectValue.String() != "CREATE TRIGGER" {
    // handles successfully trigger creation..
    return true
  }
  return false
}


// Creates an order.

func createOrder(OrderData struct{

  customerEmail string},
  destinationAddress struct{
   state string 
   city string 
   zipcode string 
   street string 
   house int
  }) (*Order){


  orderTime := time.Now()

  orders := database.Model(&Order{}).Where("HAVING id").Statement.ReflectValue
  order_name := fmt.Sprintf("Order-%s", orders)


  newOrder := Order{

  goods : Goods{Name: "", productId: "", Quantity: "",
  TotalPrice: 20000.10, currency: "usd"},
  
  customerEmail: OrderData.customerEmail,
  destinationAddress: DestinationAddress{latitude: "", longitude: "", city: "", 
  country: "", street: "", house: destinationAddress.house}, // full address to the destination.
   
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

  order := database.Model(&order).Where("id = ?", orderId).First(order)
  orderDict := Order{}
  UpdatedFields := []map[string]interface{}{}
  RestrictedFields := orderDict.getRestrictedFields()

  for _, fieldName := range orderDict.getAllFields(){

    if value, error_ := reflections.GetField(orderDict, fieldName), value != nil && error == nil &&
    hasField, error := reflections.HasField(RestrictedFields, fieldName); hasField == false{

       newValue, err := reflections.GetField(NewOrderData, fieldName)
       error := reflections.SetField(UpdatedFields, fieldName, newValue)
      }
    }
    order.Update("customerEmail", NewOrderData.customerEmail, )
    return true 
  }




