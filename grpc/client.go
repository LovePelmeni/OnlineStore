package client



import (
    "github.com/go-grpc/grpc"
)


type gRPCOrderServer struct {

    grpc_server_host string
    grpc_server_port string
}

type OrderCredentials struct {
}


func (this *gRPCOrderServer) ProcessOrder(orderData map[string]interface{}) (*OrderCredentials){
}