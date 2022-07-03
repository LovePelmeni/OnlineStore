package test_models 

import (
	"github.com/LovePelmeni/OnlineStore/OrderCheckout/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ModelSuite struct {

	Models []gorm.Model 

	OrderData map[string]string 

	DestinationAddress map[string]string 

	CustomerData map[string]string 

	Goods map[string]string 
}
func (this *ModelSuite) SetupTest(){
	
}

func (this *ModelSuite) createOrder(t *testing.T){}

func (this *ModelSuite) deleteOrder(t *testing.T){}


func (this *ModelSuite) createDestinationAddress(t *testing.T){}

func (this *ModelSuite) deleteDestinationAddress(t *testing.T){}


func (this *ModelSuite) createGoods(t *testing.T) {}

func (this *ModelSuite) deleteGoods(t *testing.T) {}