package test_emails

// import (
// 	"github.com/LovePelmeni/OnlineStore/OrderCheckout/mocks/emails"
// 	"fmt"
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// var (
// 	RejectMessage = fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
//     "has been rejected. Check you profile for more details.")
// 	AcceptMessage = fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
//     "has been accepted. Check you profile for more details.")
// )

// func getSuccessEmailMessage(
// 	customerEmail string, OrderId string) (string){
// 	return fmt.Sprintf(RejectMessage, customerEmail, OrderId)}

// func getRejectEmailMessage(
// 	customerEmail string, OrderId string) (string){
// 	return fmt.Sprintf(AcceptMessage, customerEmail, OrderId)}

// // Order Emails Test Case.
// type OrderEmailSendSuite struct {
// 	suite.Suite

// 	*require.Assertions
// 	controller *gomock.Controller

// 	customerEmail string
// 	orderId string

// 	AcceptMessage string
// 	RejectMessage string

// 	AcceptEmailController interface{}
// 	RejectEmailController interface{}
// }

// func (this *OrderEmailSendSuite) SetupTest(){
// 	this.controller = gomock.NewController(this.T())
// 	this.customerEmail = "test@gmail.com"
// 	this.orderId = "1"
// 	this.AcceptMessage = getSuccessEmailMessage(this.customerEmail, this.orderId)
// 	this.RejectMessage = getRejectEmailMessage(this.customerEmail, this.orderId)
// 	this.AcceptEmailController =  mock_emails.NewMockEmailSenderInterface(this.controller)
// 	this.RejectEmailController =  mock_emails.NewMockEmailSenderInterface(this.controller)
// }

// // checks that the accept Order handler sends Appropriate Email. Mocked Sender Controller.
// func (this *OrderEmailSendSuite) TestAcceptEmailSended(t *testing.T){

// 	defer this.controller.Finish()
// 	response, error := this.AcceptController{}.SendOrderEmailAccepted(this.customerEmail, this.orderId)
// 	assert.Equal(this.T(), response, true, "`ACCEPT EMAIL ORDER CONTROLLER`: RESPONDED WITH FAILURE.")

// 	this.AcceptEmailController.EXPECT().SendOrderEmailAccepted().Called().Times(1)
// 	this.AcceptEmailController.EXPECT().SendOrderEmailAccepted().Called().With(this.customerEmail, this.orderId)
// }

// // checks that the reject Order handler sends Appropriate Email. Mocked Sender Controller.
// func (this *OrderEmailSendSuite) TestRejectEmailSended(t *testing.T){

// 	defer this.controller.Finish()
// 	RejectController := mock_emails.NewMockEmailSenderInterface(this.controller)
// 	response, error := RejectController.SendOrderEmailRejected(this.customerEmail, this.orderId)
// 	assert.Equal(this.T(), response, true,  "`REJECT EMAIL ORDER CONTROLLER`: RESPONDED WITH FAILURE.")

// 	this.RejectEmailController.EXPECT().SendOrderEmailRejected().Called().Times(1)
// 	this.RejectEmailController.EXPECT().SendOrderEmailRejected().Called().With(this.customerEmail, this.orderId)
// }
