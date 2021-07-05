//Models/User.go
package Models

import (
	"ecomapp/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

var db infrastructure.Database

//GetAllUsers Fetch all user data
func GetCart(cart *[]Cart) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Find(cart).Error; err != nil {
		return err
	}
	return nil
}

func AddCart(cart *Cart) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Create(cart).Error; err != nil {
		return err
	}

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
