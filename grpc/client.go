package client



import (
    "github.com/go-grpc/grpc"
    "fmt"
    "emails"
    "RealTLocation/models/models"
    "RealTLocation/emails/notifications"
    "RealTLocation/"
)


var order models.Order 
var orders []models.Order 

type gRPCOrderServer struct {

    grpc_server_host string
    grpc_server_port string
}

type OrderCredentials struct {
}


func createOrder(orderData map[string]interface{}, state string, address map[string]string, goods map[string]interface{}) (string, error){

    newOrder, error := models.Order{state: state}
    models.database.Save(&newOrder)
    if error := models.database.Save(&newOrder); error != nil && error.(errors.Error){
        loggers.DebugLogger.Println("Failed to save object to Database.")
        return false, errors.New("Failed to save.")
    }
    sended, error := notifications.NotifyEmailOrderStateUpdated(newOrder.id, state, newOrder.customerEmail)
    switch sended{
        case sended && error == nil:
            return newOrder.id, nil 
        
        case error != nil && sended != true:
            models.database.rollback() // implementing rollback...
            return nil, errors.New("TransactionFailure.")
        
        default:
            return nil, errors.New("Unknown State.")
    }
}

func cancelOrder(orderId int) (bool, error){
    // After all processes made it will delete the order obj.
    order := models.database.Model(&order).Where("id = ?", orderId).First()
    updated, error := order_api.UpdateOrderState(order.id, "canceled", order.customerEmail)
    switch updated{

        case updated:
            order.Delete()
            return true, nil

        case updated != true:
            loggers.ErrorLogger.Println(
            fmt.Sprinf("Failed to Update State of the Order with ID %s to %s", order.id, "canceled"))
            return false, errors.New("Failed")
        default:
            return false, errors.New("Unknown State.")
    }
}

func (this *OrderService) ProcessOrderRequest(orderData map[string]interface{}) (*OrderCredentials){
    orderId, error := createOrder(orderData)
    switch error{
        case error:
            return ProcessOrderResponse{} // returns Not Implemented
        case error == nil: 
            decodedResponseData = json.Marshal(ioutil.ReadAll(response), Order) 
            return ProcessOrderResponse{} // returns Success.
        default:
            return ProcessOrderResponse{} // returns bad Request
    }
}
func (this *OrderService) CancelOrderRequest(orderId map[string]string) (*){

}



