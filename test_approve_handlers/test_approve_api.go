package test_approve_api

// import (
// 	"net/http"
// 	"net/url"
// 	"testing"

// 	"github.com/LovePelmeni/OnlineStore/OrderCheckout/models"
// 	"github.com/stretchr/testify/assert"
// 	"strconv"
// )
// var order models.Order


//  // Rest API Controllers Test Cases.

// func TestApproveAcceptHandler(t *testing.T){

// 	newOrder := models.Order{}
// 	models.Database.Save(newOrder)
// 	url := url.URL{Path:"http://localhost:8000/order/accept"}
// 	url.Query().Add("orderId", strconv.FormatInt(newOrder.Id, 10))
	
// 	response, httpError := http.Post(url.String(), "application/json", nil)
// 	assert.Equal(t, response.StatusCode, 200, "Response Status Code Not Equals An Appropriate One.")
// 	assert.Nil(t, httpError.Error())

// 	assert.NotNil(t, models.Database.Model(
// 	&order).Where("id = ?", 1).First(&order))

// }



// func TestApproveRejectHandler(t *testing.T){

// 	newOrder := models.Order{}
// 	models.Database.Save(newOrder)
// 	url := url.URL{Path:"http://localhost:8000/order/accept"}
// 	url.Query().Add("orderId", strconv.FormatInt(newOrder.Id, 10))

// 	response, httpError := http.Post(url.String(), "application/json", nil)
// 	assert.EqualValues(t, response.StatusCode, 200)
// 	assert.Nil(t, httpError)

// 	assert.NotNil(t, models.Database.Model(&order).Where(
// 	"id = ?", 1).First(&order))
// }