package test_kafka 


import (
	"github.com/LovePelmeni/OnlineStore/OrderCheckout/kafka"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


// Order Response Event Handler Test Case.. 

// Basically tests the part of code that is responsible to send event Back To Main Service 
// With The Decision Made About `Accept` or `Reject`

type KafkaOrderEventSenderSuite struct {
	suite.Suite 
	AcceptMessage string // message to send in `Accept` Event 
	RejectMessage string // message to send in `Reject` Event
	OrderId string // order identifier.
}

func (this *KafkaOrderEventSenderSuite) SetupTest(){
	this.AcceptMessage = ""
	this.RejectMessage = ""
	this.OrderId = "1"
}

// Mocked
func (this *KafkaOrderEventSenderSuite) TestKafkaAcceptEventHandler(t *testing.T){}

 
// Mocked
func (this *KafkaOrderEventSenderSuite) TestKafkaRejectEventHandler(t *testing.T){}