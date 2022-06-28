package models 

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"strconv"
	"time"
	"fmt"
	"reflect"
)

var (

dbHost = os.Getenv("DATABASE_HOST")
dbUser = os.Getenv("DATABASE_USER")
dbPassword = os.Getenv("DATABASE_PASSWORD")
dbPort = os.Getenv("DATABASE_PORT")
dbName = os.Getenv("DATABASE_NAME")

database, error = gorm.Open(postgres.New{
    DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
    dbHost, dbUser, dbPassword, dbName, dbPort),
})
)

type CustomerCredentials struct {
    email string `gorm:"VARCHAR(50) NOT NULL"`
    phone_number string `gorm:"VARCHAR(50) NOT NULL"`
}
func (this *CustomerCredentials) validateCustomerCredentials() (bool, error){
	return false, nil
}

type DestinationAddress struct {
	gorm.Model
	ID int64
	latitude string `gorm:"VARHCHAR(50) NOT NULL`
	longitude string `gorm:"VARHCHAR(50) NOT NULL`
	street string `gorm:"VARHCHAR(50) NOT NULL`
	house string `gorm:"VARHCHAR(50) NOT NULL`
	city string `gorm:"VARHCHAR(50) NOT NULL`
	country string `gorm:"VARHCHAR(50) NOT NULL`
}
func (this *DestinationAddress) validateLocation() (bool, error){
	// validates the Destination before saving...
}

type Order struct {
	gorm.Model 
	ID int64
	order_name string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	goods Goods `gorm:"foreignKey Goods NOT NULL"`
	createdAt time.Time `gorm:"DATE DEFAULT CURRENT_DATE NOT NULL"`
	customerCredentials CustomerCredentials `gorm:"foreignKey CustomerCredentials NOT NULL"`
	destinationAddress DestinationAddress `gorm:"foreignKey DestinationAddress NOT NULL"`
	customerEmail string `gorm:"VARCHAR(50) NOT NULL`
}
func (this *Order) validateOrder() (bool, error){
	// validates Order Info before saving the object....
}

type Goods struct {
	gorm.Model
	ID int64
	Name string `gorm:"VARCHAR(50) NOT NULL"`
	Quantity string `gorm:"INTEGER NOT NULL"`
	TotalPrice float64 `gorm:"NUMERIC(10, 2) NOT NULL`
}

func (this *Goods) validateGoods() (bool, error){
	// validates goods info before saving...
}


