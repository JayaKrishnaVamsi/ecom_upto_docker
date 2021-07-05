//Controllers/User.go
package Controllers

import (
	"ecomapp/Models"
	"ecomapp/infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetPayment(c *gin.Context) {
	var payment []Models.Payment
	err := Models.GetPayment(&payment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

//CreateUser ... Create User
func CreatePayment(c *gin.Context) {
	var payment Models.Payment
	//var cart Models.Cart
	c.BindJSON(&payment)
	err := Models.CreatePayment(&payment)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
	fmt.Println("before")
	DoPayment(&payment)
}

var db infrastructure.Database

func DoPayment(payment *Models.Payment) {
	db = infrastructure.NewDatabase()
	fmt.Println("after")
	po := payment.OId
	var result []Models.Cart
	db.DB.Raw("SELECT * FROM cart WHERE order_id = ?", po).Scan(&result)
	fmt.Println(result)
	var total int = 0
	var r1 Models.Product
	for _, v := range result {
		db.DB.Raw("SELECT * FROM product WHERE id = ?", v.PId).Scan(&r1)
		total = total + (int(r1.Price) * int(v.PQty))
	}
	fmt.Println("Total Bill")
	fmt.Println(total)
	db.DB.Exec("UPDATE payment SET bill = ? WHERE o_id = ?", total, po)
	fmt.Println("Updated bill in db")
	var qt uint
	db.DB.Raw("SELECT quantity FROM product WHERE id = ?", result[0].PId).Scan(&qt)

	for _, v := range result {
		qt = qt - v.PQty
		db.DB.Exec("UPDATE product SET quantity = ? WHERE id = ?", qt, v.PId)

	}
	fmt.Println("Quantity decremented in db")
}

/*
//GetUserByID ... Get the user by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateUser ... Update the user information
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
*/
//DeleteUser ... Delete the user
/*func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}*/
