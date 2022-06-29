package tests

import (
	"assert"
	"net/http"
	"testing"
	"url"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
	"json"
	"RealTLocation/models/models"
	"RealTLocation/services/order_api"
)

var customer models.Customer 
var order models.Order 

type MockedKafkaEvent struct {
	event kafka.Event 
}

func getMockedKafkaEvent() (kafka.Event){
	event := MockedKafkaEvent{event: kafka.Event()}
	return event.event 
}

type MockedgRPCRequest struct {
	request http.Request
	context grpc.DialContext
}

type MockedOrderData struct {
	// mocked order
}

func getMockedgRPCRequest(request_params map[string]interface{},
request_data map[string]interface{}) (http.Request){
	// returns prepared Http Request For Mocked GRPC Event..
	req_url := url.URL("http://some-test-url")
	req_url.queryParams = request_params 
	request := http.Request{URL: req_url}
	request.Body = []byte(json.Marshal(request_data))
	return request
}

func getMockedOrderData() (MockedOrderData){
	return MockedOrderData{}
}

func mockedKafkaResponseHandler() (){
	//
}


// func TestEventResponseHandlerCase(t *testing.T){
// 	mocked_handler := mockedKafkaResponseHandler()()
// 	mocked_test_valid_handler_data := func() (kafka.Event){
// 		return kafka.Event{""}
// 	}()
// 	mocked_test_invalid_handler_data := func() (map[string]string){
// 		return map[string]string{"invalid_data": "invalid_content"}
// 	}()
// 	response, error := mocked_handler(mocked_test_valid_handler_data)
// 	assert.True(response)
// 	assert.Empty(error)

// 	invalid_response, error := mocked_handler(mocked_test_invalid_handler_data)
// 	assert.AnError(error)
// 	assert.Empty(invalid_response)
// }

func TestOrderCreateCase(t *testing.T){ // mocked 


	mockedOrderData := getMockedOrderData()
	query_params := map[string]string{"orderId": "", "customerId": ""}
	mocked_grpc_request := getMockedgRPCRequest(query_params, mockedOrderData)
	
	method := func(request http.Request) (bool, error){
		if response, error := order_api.ProcessOrderRequest(); error != nil {
			return false, error 
		} else {
			return true, error 
		}
	}
	response, error := method(mocked_grpc_request)
	assert.Empty(error)
	assert.NotEmpty(response)

}

func TestOrderUpdateCase(t *testing.T){ // mocked 

	query_params := map[string]int{"orderId": 1, "customerId": 2}
	mocked_grpc_request := getMockedgRPCRequest(query_params, mocked)
	
	method := func(request http.Request) (bool, error){
		if response, error := order_api.ProcessOrderRequest(); error != nil {
			return false, error 
		} else {
			return true, error 
		}
	}
	response, error := method(mocked_grpc_request)
	assert.Empty(error)
	assert.NotEmpty(response)
}

func TestOrderDeleteCase(t *testing.T){  // mocked
	customer := models.Customer{}
	models.database.Save(&customer)
	order := models.Order{}
	models.database.Save(&order)

	query_params := map[string]int{"orderId": 1, "customerId": 2}
	mocked_grpc_request := getMockedgRPCRequest(query_params, req_data)
	
	method := func(request http.Request) (bool, error){
		if response, error := order_api.ProcessOrderRequest(); error != nil {
			return false, error 
		} else {
			return true, error 
		}
	}
	response, error := method(mocked_grpc_request)
	assert.Empty(error)
	assert.NotEmpty(response)
}



func TestOrdersGetCase(t *testing.T){
	urlAddress := url.URL("http://localhost:8000/delivered/orders/")
	urlAddress.queryParams[""] = "TestEmail@gmail.com"
	response, error := http.Get(urlAddress.String())
	assert.Equals(response.StatusCode, 200)
	assert.Empty(error)
}
