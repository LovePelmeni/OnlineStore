package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	_ "strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (

dbHost = os.Getenv("DATABASE_HOST")
dbUser = os.Getenv("DATABASE_USER")
dbPassword = os.Getenv("DATABASE_PASSWORD")
dbPort = os.Getenv("DATABASE_PORT")
dbName = os.Getenv("DATABASE_NAME")


databaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
Database, databaseError = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
	DisableForeignKeyConstraintWhenMigrating: true,
}))


var (
  DebugLogger *log.Logger 
  ErrorLogger *log.Logger 
  WarningLogger *log.Logger 
  InfoLogger *log.Logger 
)

var (
  OrderStates = []string{"Pre-Confirmed", "Confirmed", "Canceled", "Delivered"}
)
var order *Order 


type RestrictedFields interface {
    // Interface allows to return fields that is not allowed to be updated.
    GetRestrictedFields() ([]string)
}


type DestinationAddress struct {
  // structure reprensents the destination Address
	Id int 
	Latitude string `gorm:"VARCHAR(100) NOT NULL DEFAULT '0.0000'"`
	Longitude string `gorm:"VARCHAR(100) NOT NULL DEFAULT '0.0000'"`

  City string `gorm:"VARCHAR(100) NOT NULL"`
  Country string `gorm:"VARCHAR(100) NOT NULL"`

	Street string `gorm:"VARCHAR(100) NOT NULL"`
	House int `gorm:"INTEGER NOT NULL"` 
}


type Goods struct {
  // structure represents the info about the delivery product
	Id int 
	Name string  `gorm:"VARCHAR(40) NOT NULL"`// names of the products that has been purchased..
  ProductId string `gorm:"VARCHAR(40) NOT NULL"` // identifiers of the products that has been purchased...
	Quantity string `gorm:"VARCHAR(50) NOT NULL DEFAULT '1'"`
	TotalPrice string `gorm:"NUMERIC(10, 2) NOT NULL"`
  Currency string `gorm:"VARCHAR(10) NOT NULL DEFAULT 'usd'"`
}

func (this *Goods) createGoodsObject(GoodsInfoData map[string]string) (*Goods, error){
  newGoodsObj := Goods{

        Name: GoodsInfoData["Name"],
        ProductId: GoodsInfoData["ProductId"],
        Quantity: GoodsInfoData["Quantity"],
        TotalPrice: GoodsInfoData["TotalPrice"],
        Currency: GoodsInfoData["Currency"],
  }
  if gormGoodsObject := Database.Table("goods").Create(&newGoodsObj);
   gormGoodsObject != nil{
    return &newGoodsObj, nil
  }
  return nil, errors.New(
  "Failed to Create Goods ORM Object.")
}


type Order struct {
  gorm.Model
	Id int 
	Order_name string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`

  DestinationAddressId int  // reference identifier to the DestinationAddress table..
  OrderedGoodsId int //reference identifier to the Goods Table..

	OrderedGoods Goods `gorm:"foreignKey:OrderedGoodsId;references:Id; constraint:OnUpdate:PROTECT,OnDelete:CASCADE;"` // 
	CreatedAt time.Time `gorm:"DATE DEFAULT CURRENT_DATE"`
	Destination DestinationAddress `gorm:"foreignKey:DestinationAddressId;references:Id; constraint:OnUpdate:PROTECT,OnDelete:CASCADE;"`
	CustomerEmail string `gorm:"VARCHAR(100) NOT NULL"`
	State string `gorm:"VARCHAR(10) NOT NULL"`
}


func (this *Order) getRestrictedFields() ([]string){
  // returns restricted fields for Order Model..
  return []string{"CreatedAt", 
  "DestinationAddress", "Order_name", "Goods", "Id"}
}

type OrderCustomerCredentials struct {
    // Structure represents the
    CustomerEmail string
    PhoneNumber string
}


func CreateRestrictColumnUpdateTrigger(
  // this Method Adds Restrictions for fields to Update.. 
  // Receives the model and list of fields which should not be updated.
model *gorm.DB, forbidden_columns []string) (bool){
  sql_first_part_command := `CREATE OR REPLACE FUNCTION restrictUpdateField RETURNS TRIGGER 
  LANGUAGE plpgsql AS
  $$BEGIN` 
  value := forbidden_columns 
  sql_first_part_command += fmt.Sprintf(`IF NEW.%s IS DISTINCT FROM OLD.%s` + 
  `NEW.%s = OLD.%s`, value, value, value, value)

  sql_first_part_command += "END IF RETURN NEW; END;$$;"
  sql_second_part_command := fmt.Sprintf("CREATE TRIGGER update_restrictor BEFORE UPDATE %s" + 
  "FOR EACH ROW EXECUTE PROCEDURE restrictUpdateField", model.Name())
  error := model.Exec(
  sql_first_part_command + sql_second_part_command)
  if error.Statement.ReflectValue.String() != "CREATE TRIGGER" {
    // handles successfully trigger creation..
    return true
  }
  return false
}

// Creates an order.

func CreateOrder(customerCredentials OrderCustomerCredentials,
  destinationAddress DestinationAddress, goods Goods) (*Order){
  
  orderTime := time.Now()
  order_name := fmt.Sprintf("Order-%s", orderTime)

  newOrder := Order{

        OrderedGoods : goods,
    
        CustomerEmail: customerCredentials.CustomerEmail,
        Destination: destinationAddress, // full address to the destination.
        
        Order_name: order_name,

        CreatedAt: orderTime,
        State: fmt.Sprintf("%s", OrderStates[0]), // current state of the order. default "pre-confirmed"
  }
  Database.Save(&newOrder)
  return &newOrder
}



// Deletes order
func DeleteOrder(orderId string) (bool){
  Database.Model(&Order{}).Where("id = ?", orderId).Delete(order)
  return true 
}


func UpdateOrderState(orderId string, State string) (bool){

   lockedTable := Database.Clauses(clause.Locking{Strength: "OnUpdate",
   Table: clause.Table{Name: "Order"}})
   lockedTable.AllowGlobalUpdate = false 

  order := Database.Model(order).Where(
  "id = ?", orderId).Update("State", State)
  DebugLogger.Println(fmt.Sprintf("State of Order ID: %s has been updated to %s",
  orderId, State))
  defer lockedTable.Commit()
  if order.Error != nil {return false} else {return true}
}
