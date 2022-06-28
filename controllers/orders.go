package orders

import (
	"RealTLocation/models/models"
	"json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var order models.Order

func GetDeliveredOrders(context *gin.Context) {
	customerEmail := context.Request.URL.Query().Get("customerEmail")
	filteredOrders := models.database.Model(
		&order).Where("customerCredentials.email = ?", customerEmail).All()
	context.JSON(http.StatusOK, gin.H{"orders": json.Marshal(filteredOrders)})
}

