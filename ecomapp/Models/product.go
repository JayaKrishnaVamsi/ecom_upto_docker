//Models/User.go
package Models

import (
	"ecomapp/infrastructure"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db2 infrastructure.Database

//GetAllUsers Fetch all user data
func GetAllProducts(product *[]Product) (err error) {
	db2 = infrastructure.NewDatabase()
	if err = db2.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateProduct(product *Product) (err error) {
	db2 = infrastructure.NewDatabase()
	if err = db2.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetProductByID(product *Product, id string) (err error) {
	db2 = infrastructure.NewDatabase()
	if err = db2.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(product *Product, id string) (err error) {
	db2 = infrastructure.NewDatabase()
	fmt.Println(product)
	db2.DB.Save(product)
	return nil
}

//DeleteUser ... Delete user
func DeleteProduct(product *Product, id string) (err error) {
	db2 = infrastructure.NewDatabase()
	db2.DB.Where("id = ?", id).Delete(product)
	return nil
}
