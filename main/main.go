package main

import (
	"customers"
	"middlewares"
	"os"
	"permissions"
	"products"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"net/http"
	"github.com/gorilla/csrf"
)

var (
	applicationHost = os.Getenv("APPLICATION_HOST")
	applicationPort = os.Getenv("APPLICATION_PORT")
	paymentApplicationHost = os.Getenv("PAYMENT_APPLICATION_HOST")
)

func main() {

	DEBUG := os.Getenv("DEBUG")

	if DEBUG {
		DATABASE_NAME := os.Getenv("DATABASE_NAME")
		DATABASE_HOST := os.Getenv("DATABASE_HOST")
		DATABASE_USER := os.Getenv("DATABASE_USER")
		DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
		DATABASE_PORT := os.Getenv("DATABASE_PORT")
	} else {
		DATABASE_NAME := os.Getenv("DATABASE_NAME")
		DATABASE_HOST := os.Getenv("DATABASE_HOST")
		DATABASE_USER := os.Getenv("DATABASE_USER")
		DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
		DATABASE_PORT := os.Getenv("DATABASE_PORT")
	}

	csrf_byte_string := `KSazoE72qxuMUcOmBw7Ye3DClxdRCTVr9ZL0msfhV7r3s7DkIQdqPYx5U49k4iGl`
	router := gin.Default()
	credentials := handlers.AllowCredentials()
	allowed_methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"})
	allowed_origins := handlers.AllowedOrigins([]string{"http:" + applicationHost + ":8000", paymentApplicationHost})
	allowed_headers := handlers.AllowedHeaders([]string{"*"})
	csrf.Protect([]byte(csrf_byte_string))
	

	router.Group("customer/")
	{

		router.POST("/create/customer/", customers.CreateCustomerController)
		router.PUT("/update/customer/", middlewares.AuthenticationMiddleware(), customers.UpdateCustomerController)
		router.DELETE("delete/customer/", middlewares.AuthenticationMiddleware(), customers.DeleteCustomerController)
		router.GET("get/customer/", middlewares.AuthenticationMiddleware(), customers.GetCustomerProfileController)

	}

	router.Group("product/").Use(middlewares.AuthenticationMiddleware)
	{

		router.GET("get/all/products/", products.GetAllProductsController)
		router.POST("create/product/", permissions.IsProductOwner(), products.CreateProductController)
		router.PUT("update/product/", permissions.IsProductOwner(), products.UpdateProductController)
		router.DELETE("delete/product/", permissions.IsProductOwner(), products.DeleteProductController)

	}
	// liked Product API Controllers.
	router.Group("liked/product/").Use(middlewares.AuthenticationMiddleware)
	{

		router.POST("add/liked/product", products.AddLikedProducts)
		router.DELETE("remove/liked/product", products.RemoveLikedProduct)

	}
	http.ListenAndServe(applicationHost + ":" + applicationPort, handlers.CORS(
	credentials, allowed_methods, allowed_headers, allowed_origins)(router))
	router.Run(applicationHost + ":" + applicationPort)
}




