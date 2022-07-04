package test_models

// import (
// 	"testing"
// 	"gorm.io/gorm"
// 	"github.com/LovePelmeni/OnlineStore/OrderCheckout/models"
// 	"github.com/stretchr/testify/assert"
// 	"os"
// )

// type TestDatabaseConnection struct {
// 	DatabaseHost string 
// 	DatabasePort string 
// 	DatabaseName string 
// 	DatabaseUser string 
// 	DatabasePassword string 
// }

// func (this *TestDatabaseConnection) databaseConnection() (*gorm.DB, error){

// }

// type ModelSuite struct {

// 	DatabaseConnection *TestDatabaseConnection

// 	Models []struct{}

// 	OrderData map[string]string

// 	DestinationAddress map[string]string

// 	CustomerData map[string]string

// 	Goods map[string]string
// }

// func (this *ModelSuite) SetupTest() {

// 	this.DatabaseConnection = &TestDatabaseConnection{
// 		DatabaseHost: os.Getenv("TEST_DATABASE_HOST"),
// 		DatabasePort: os.Getenv("TEST_DATABASE_PORT"),
// 		DatabaseUser: os.Getenv("TEST_DATABASE_USER"),
// 		DatabasePassword: os.Getenv("TEST_DATABASE_PASSWORD"),
// 		DatabaseName: os.Getenv("TEST_DATABASE_NAME"),
// 	} 
// 	this.Models = []struct{}{models.Order,
// 	models.DestinationAddress, models.Goods}

// 	this.OrderData = map[string]string{}
// 	this.DestinationAddress = map[string]string{}
// 	this.CustomerData = map[string]string{}
// 	this.Goods = map[string]string{}
// }


// func (this *ModelSuite) TestModelCreate(t *testing.T) {
// 	connection, error_ := this.DatabaseConnection.databaseConnection()
// 	if error_ != nil {assert.Error(t, error_, "Test Database Connection  Failed.")}
// 	for _, model := range this.Models{
// 		if created := connection.Model(&model).Create(); created.Error() != nil {
// 			assert.Error(t, created.Error(), "Model Creation Failed.")
// 		}
// 	}
// }

// func (this *ModelSuite) TestModelUpdate(t *testing.T) {

// }

// func (this *ModelSuite) TestModelDelete(t *testing.T){

// }