//Models/User.go
package Models

import (
	"ecomapp/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

var db1 infrastructure.Database

//GetAllUsers Fetch all user data
func GetPayment(payment *[]Payment) (err error) {
	db1 = infrastructure.NewDatabase()
	if err = db1.DB.Find(payment).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreatePayment(payment *Payment) (err error) {
	db1 = infrastructure.NewDatabase()
	if err = db1.DB.Create(payment).Error; err != nil {
		return err
	}
	//db.DB.Exec("UPDATE product SET quantity = ? WHERE id = ?", gorm.Expr("quantity - ?", cart.PQty), cart.PId)
	return nil
}

//GetUserByID ... Fetch only one user by Id
/*func GetProductByID(product *Product, id string) (err error) {
	if err = db.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}*/

//UpdateUser ... Update user
/*func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	db.DB.Save(product)
	return nil
}*/

//DeleteUser ... Delete user
/*func DeleteProduct(product *Product, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(product)
	return nil
}
*/
