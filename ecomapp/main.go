package main

import (
	"ecomapp/Controllers"
	"ecomapp/Models"
	"ecomapp/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//Config.DB.

	infrastructure.LoadEnv()
	db := infrastructure.NewDatabase()
	db.DB.AutoMigrate(&Models.Customer{})
	db.DB.AutoMigrate(&Models.Product{})
	db.DB.AutoMigrate(&Models.Cart{})
	db.DB.AutoMigrate(&Models.Payment{})
	router.GET("/", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})

	})
	grp1 := router.Group("/customer")
	{
		grp1.GET("all_customers", Controllers.GetCustomers)
		grp1.POST("customers", Controllers.CreateCustomer)
		grp1.GET("customers/:id", Controllers.GetCustomerByID)
		grp1.PUT("customers/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customers/:id", Controllers.DeleteCustomer)
	}
	grp2 := router.Group("/product")
	{
		grp2.GET("all_products", Controllers.GetProducts)
		grp2.GET("products/:id", Controllers.GetProductByID)

	}
	grp3 := router.Group("/mycart")
	{
		grp3.GET("allitems", Controllers.GetCart)
		grp3.POST("additem", Controllers.AddCart)

	}
	grp4 := router.Group("/payment")
	{
		grp4.GET("allpayments", Controllers.GetPayment)
		grp4.POST("pay", Controllers.CreatePayment)

	}
	router.Run(":8000")
}
