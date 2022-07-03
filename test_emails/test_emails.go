package test_emails 



import (
	"github.com/LovePelmeni/OnlineStore/OrderCheckout/emails"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


func TestSendEmail(){

}

var (
	RejectMessage = fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
    "has been rejected. Check you profile for more details.")
	AcceptMessage = fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
    "has been accepted. Check you profile for more details.")
)


func getSuccessEmailMessage(
	customerEmail string, OrderId string) (string){
	return fmt.Sprintf(RejectMessage, customerEmail, OrderId)}
	
	
func getRejectEmailMessage(
	customerEmail string, OrderId string) (string){
	return fmt.Sprintf(AcceptMessage, customerEmail, OrderId)}
	
	


// Order Emails Test Case.
type OrderEmailSendSuite struct {
	suite.Suite 
	customerEmail string 
	orderId string 
	AcceptMessage string 
	RejectMessage string 
}

func (this *OrderEmailSendSuite) SetupTest(){
	this.customerEmail = "test@gmail.com"
	this.orderId = "1"
	this.AcceptMessage = getSuccessEmailMessage(this.customerEmail, this.OrderId)
	this.RejectMessage = getRejectEmailMessage(this.customerEmail, this.OrderId)
}

// checks that the accept Order handler sends Appropriate Email. Mocked Sender Controller.
func (this *OrderEmailSendSuite) TestAcceptEmailSended(t *testing.T){}


// checks that the reject Order handler sends Appropriate Email. Mocked Sender Controller.
func (this *OrderEmailSendSuite) TestRejectEmailSended(t *testing.T){}