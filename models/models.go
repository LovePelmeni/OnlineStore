package models

import (
	"fmt"
	"os"
	_ "reflect"
	_ "strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = os.Getenv("DATABASE_HOST")
	dbUser     = os.Getenv("DATABASE_USER")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
	dbPort     = os.Getenv("DATABASE_PORT")
	dbName     = os.Getenv("DATABASE_NAME")

	databaseURL     = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	database, error = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
)

type DestinationAddress struct {
	id        int64
	latitude  string `gorm:"VARHCHAR(50) NOT NULL"`
	longitude string `gorm:"VARHCHAR(50) NOT NULL"`
	street    string `gorm:"VARHCHAR(50) NOT NULL"`
	house     string `gorm:"VARHCHAR(50) NOT NULL"`
	city      string `gorm:"VARHCHAR(50) NOT NULL"`
	country   string `gorm:"VARHCHAR(50) NOT NULL"`
}

type OrderState struct {
	order_state string // represents current order state...
}

type Order struct {
	id                 int64
	order_name         string             `gorm:"VARCHAR(50) NOT NULL UNIQUE"`
	goods              Goods              `gorm:"foreignKey Goods NOT NULL"`
	createdAt          time.Time          `gorm:"DATE DEFAULT CURRENT_DATE NOT NULL"`
	destinationAddress DestinationAddress `gorm:"foreignKey DestinationAddress NOT NULL"`
	customerEmail      string             `gorm:"VARCHAR(50) NOT NULL"`
	state              string             `gorm:"VARCHAR(10) NOT NULL"`
}

type Goods struct {
	id         int64
	Name       string  `gorm:"VARCHAR(50) NOT NULL"`
	Quantity   string  `gorm:"INTEGER NOT NULL"`
	TotalPrice float64 `gorm:"NUMERIC(10, 2) NOT NULL"`
}
