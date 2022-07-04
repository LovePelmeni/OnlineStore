package test_kafka 


// import (
// 	"github.com/LovePelmeni/OnlineStore/OrderCheckout/mocks/kafka"
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// // Order Response Event Handler Test Case.. 

// // Basically tests the part of code that is responsible to send event Back To Main Service 
// // With The Decision Made About `Accept` or `Reject`

// func getKafkaAcceptMessage() (string){ // returns Acceptance Event message for Kafka Event.
// 	return ""
// }

// func getKafkaRejectMessage() (string){ // returns Reject Kafka Event Message for Kafka Event
// 	return ""
// }

// type KafkaOrderEventSenderSuite struct {
// 	suite.Suite 
// 	require.Assertions
// 	controller *gomock.Controller
// 	AcceptMessage string // message to send in `Accept` Event 
// 	RejectMessage string // message to send in `Reject` Event
// 	OrderId int64// order identifier.
// 	AcceptMessageController interface{}// mocked Accept Handler Controller 
// 	RejectMessageController interface{}
// }


// func (this *KafkaOrderEventSenderSuite) SetupTest(){
// 	// Set Up
// 	this.controller = gomock.NewController(this.T())
// 	this.AcceptMessage = getKafkaAcceptMessage()
// 	this.RejectMessage = getKafkaRejectMessage()
// 	this.OrderId = []int64{1}[0]
// 	this.AcceptMessageController = mock_kafka.TestKafkaAcceptEventHandler(this.controller)
// 	this.RejectMessageController = mock_kafka.TestKafkaRejectEventHandler(this.controller)
// }


// // Tests Acceptance Event Handlers for Events...
// func (this *KafkaOrderEventSenderSuite) TestKafkaAcceptEventHandler(t *testing.T){

// 	defer this.controller.Finish()

// 	this.AcceptMessageController.EXPECT().SendAcceptOrderEvent().Called().Times(1).Return(200)
// 	this.AcceptMessageController.EXPECT().SendAcceptOrderEvent().Called().With(this.AcceptMessage, this.OrderId).Return(200)
// }
 
// // Tests Reject Event Handlers for Events...
// func (this *KafkaOrderEventSenderSuite) TestKafkaRejectEventHandler(t *testing.T){

	
// 	defer this.controller.Finish()
// 	this.RejectMessageController.EXPECT().SendRejectOrderEvent().Called().Times(1).Return(200)
// 	this.RejectMessageController.EXPECT().SendRejectOrderEvent().Called().With(this.RejectMessage, this.OrderId).Return(200)
// 	this.RejectMessageController.EXPECT().SendRejectOrderEvent(gomock.Eq(true, nil)).Return(200)
// }


