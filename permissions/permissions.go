package permissions

import (
	"logging"
	"models"

	"github.com/gin-gonic/gin"
)

var product models.Product
var products []models.Product
var customer models.Customer
var customers []models.Customer

type PermissionDenied struct {
	exception error
}

func (product *models.Product) CheckIsOwner(customer models.Customer) bool {
	if product.owner != customer {
		return false
	}
	return true
}

func IsProductOwnerPermission() gin.HandlerFunc {

	return func(context *gin.Context) {
		customer, error = models.database.Model(&customer).Where("id = ?", context.Query("customer_id"))
		product_id := context.Query("product_id")
		product, error := models.database.Model(&product).Where("id = ?", product_id).First()
		if error != nil || product == nil {
			logging.DebugLogger.Println("Exception while obtaining product.")
			return
		}
		if is_owner, error := product.CheckIsOwner(customer); is_owner {
			context.Next()
		}
		return
	}
}
