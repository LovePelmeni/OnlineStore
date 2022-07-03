package test_approve_api

// import (
// 	"net/http"
// 	"net/url"
// 	"testing"

// 	"github.com/LovePelmeni/OnlineStore/OrderCheckout/models/"
// 	"github.com/stretchr/testify/assert"
// )

// var order models.Order 

// func TestApproveAcceptHandler(t *testing.T) {
// 	newOrder := models.Order{}
// 	models.Database.Save(newOrder)
// 	url := url.URL{Path:"http://localhost:8000/order/accept"}
// 	url.Query().Add("orderId", newOrder.Id)
// 	response, httpError := http.Post(url.String(), "application/json", nil)
// 	assert.Equal(response.StatusCode, 200)
// 	assert.IsEmpty(httpError)
// 	assert.IsNotEmpty(models.Database.Model(models.Order).Where("id = ?", 1).First(order))
// }

// func TestApproveRejectHandler(t *testing.T){

// 	newOrder := models.Order{}
// 	models.Database.Save(newOrder)
// 	url := url.URL{Path:"http://localhost:8000/order/accept"}
// 	url.Query().Add("orderId", newOrder.Id)
// 	response, httpError := http.Post(url.String(), "application/json", nil)
// 	assert.EqualValues(response.StatusCode, 200)
// 	assert.IsEmpty(httpError)
// 	assert.IsNotEmpty(models.Database.Model(models.Order).Where("id = ?", 1).First(order))
// }