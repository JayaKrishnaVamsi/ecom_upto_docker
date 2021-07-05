//Models/Customer.go
package Models

import (
	"fmt"

	"ecomapp/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

var db3 infrastructure.Database

//GetAllCustomers Fetch all user data
func GetAllCustomers(customer *[]Customer) (err error) {
	db3 = infrastructure.NewDatabase()
	if err = db3.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

//CreateCustomer ... Insert New data
func CreateCustomer(customer *Customer) (err error) {
	db3 = infrastructure.NewDatabase()
	if err = db3.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetCustomerByID ... Fetch only one user by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	db3 = infrastructure.NewDatabase()
	if err = db3.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCustomer ... Update user
func UpdateCustomer(customer *Customer, id string) (err error) {
	db3 = infrastructure.NewDatabase()
	fmt.Println(customer)
	db3.DB.Save(customer)
	return nil
}

//DeleteCustomer ... Delete user
func DeleteCustomer(customer *Customer, id string) (err error) {
	db3 = infrastructure.NewDatabase()
	db3.DB.Where("id = ?", id).Delete(customer)
	return nil
}
