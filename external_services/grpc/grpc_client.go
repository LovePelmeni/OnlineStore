package grpc_client

import (
	"google.golang.org/grpc"
)

type BasegRPCClient struct{}

var (
	serviceNames = map[string]string{"intent": "PaymentIntent", "session": "PaymentSession"}
)

func (this *BasegRPCClient) createClient(service_name string) grpc.ClientConn {
	return grpc.ClientConn()
}

func (this *BasegRPCClient) sendgRPCRequest(client struct{}) bool {
	return grpc.ClientConn()
}

type PaymentClient struct {
	// model of the gRPC Client responsible for delivery messages related to Payments.

}

func createPaymentIntentClient() grpc.ClientConn {
	service_name := serviceNames["intent"]
	client := BasegRPCClient{}
	paymentIntentClient := client.createClient(service_name)
	return paymentIntentClient
}

func createPaymentSessionClient() grpc.ClientConn {
	return grpc.ClientConn()
}

func (this *PaymentClient) createPaidPaymentClient() grpc.ClientConn {
	return grpc.ClientConn()
}

func SendGRPCRequest() bool {
	return true
}

type CheckoutClient struct {
	// model of the gRPC Client resposible for delivery messages related to "Checkouts"

}

func (this *CheckoutClient) createCheckoutClient() grpc.ClientConn {
	return grpc.ClientConn()
}
func (this *CheckoutClient) sendGetPaymentCheckoutRequest() grpc.ClientConn {
	return grpc.ClientConn()
}
