package models 

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"strconv"
	"time"
)

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
func (this *DestinationAddress) validateLocation(){
	// validates the Destination before saving...

}

type Order struct {
	gorm.Model 
	ID int64
	order_name string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	goods Goods `gorm:"foreignKey Goods NOT NULL"`
	createdAt time.Time `gorm:"DATE DEFAULT CURRENT_DATE NOT NULL"`
	destinationAddress DestinationAddress `gorm:"foreingKey DestinationAddress NOT NULL"`
	customerEmail string `gorm:"VARCHAR(50) NOT NULL`
}
func (this *Order) validate_order(){
	// validates Order Info before saving the object....

}

type Goods struct {
	gorm.Model
	ID int64
	Name string `gorm:"VARCHAR(50) NOT NULL"`
	Quantity string `gorm::"INTEGER NOT NULL"`
	TotalPrice float64 `gorm:"NUMERIC(10, 2) NOT NULL`
}
func (this *Goods) validate_goods(){
	// validates goods info before saving...
}