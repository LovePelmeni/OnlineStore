package test_grpc_response_handlers 



import (
	"testing"
	"github.com/go-mock/"
	"json"
	"grpc_client"
)

type TestGRPCResponse struct {

	response_content []json.FieldQueryString  
	statusCode int64 
}

func MockedPaymentIntentResponse() TestGRPCResponse {
	response_content := json.Marshal{"user": "customer"}
	return TestGRPCResponse{
		response_content: response_content,
		statusCode: 200,
	}
}

func TestGRPCPaymentResponseHandler(t * testing.T){
	//
	response := MockedPaymentIntentResponse()
	// 
}