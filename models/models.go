package main 

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"fmt"
	"main"
)

var (
	database, error = gorm.Open(postgres.New{
	DSN: fmt.Sprintf("host=%s user=%s port=%s dbname=%s port=%s", main.DATABASE_HOST, main.DATABASE_USER,
	main.DATABASE_PASSWORD, main.DATABASE_NAME, main.DATABASE_PORT)},
	&gorm.Config{})
)

type Customer struct {
	gorm.Model
	// Model Represents Customer Object 
	Username string `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	Email string `gorm:"VARCHAR(100) NOT NULL UNIQUE"`
	Password string `gorm:"VARCHAR(50) NOT NULL"`
	purchasedProducts Product `gorm:"foreignKey Product DEFAULT NULL"`
	createdAt time.Time `gorm:"DATE NOT NULL DEFAULT CURRENT_DATE"`

}

type Curreny struct {
	gorm.Model 
	// Model Represents Currency Object.
	currency string `gorm:"VARCHAR(10) NOT NULL UNIQUE`
}

type Product struct {
	gorm.Model 
	// Model Represents Product Object..
	ProductName string `gorm:"VARCHAR(50) NOT NULL"`
	ProductDescription string `gorm:"VARCHAR(400) NOT NULL`
	ProductPrice float64 `gorm:"NUMERIC(10, 5) NOT NULL`
	ProductCurrency string `gorm:"foreignKey Currency REFERENCES Currency.currency NOT NULL"`

}