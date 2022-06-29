package client



import (
    "google.golang.org/grpc"
    "fmt"
    "RealTLocation/models/models"
    "RealTLocation/emails/notifications"
    "RealTLocation/"
    "errors"
    "RealTLocation/loggers/loggers"
    "RealTLocation/services/order_api"
    "json"
    "ioutil"
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

    newOrder := models.Order{state: state}
    models.database.Save(&newOrder)

    sended, error := notifications.NotifyEmailOrderStateUpdated(newOrder.id, state, newOrder.customerEmail)
    switch sended{
        case sended && error == nil:
            return newOrder.id, nil 
        
        case error != nil && sended != true:
            models.database.rollback() // implementing rollback...
            return "", errors.New("TransactionFailure.")
        
        default:
            return "", errors.New("Unknown State.")
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
            fmt.Sprintf("Failed to Update State of the Order with ID %s to %s", order.id, "canceled"))
            return false, errors.New("Failed")
        default:
            return false, errors.New("Unknown State.")
    }
}

func (this *OrderService) ProcessOrderRequest(orderData map[string]interface{}) (*OrderCredentials){
    
    decodedOrderData := json.Unmarshal(ioutil.ReadAll(orderData))
    orderId, error_ := createOrder(orderData, 
    decodedOrderData["state"], decodedOrderData["address"], decodedOrderData["goods"])

    switch error_{
        case error_: 
            return ProcessOrderResponse{orderId: nil} // returns BadRequest
        default:
            return ProcessOrderResponse{orderId: orderId} // returns Success.
    }
}
func (this *OrderService) CancelOrderRequest(orderId map[string]string) (*OrderCanceledResponse){
    decodedId := orderId["orderId"]
    canceled, error := cancelOrder(decodedId)
    if canceled != true || error != nil {return OrderCanceledResponse{
    status: "Failed"}}else{return OrderCanceledResponse{status: "Success"}}
}
