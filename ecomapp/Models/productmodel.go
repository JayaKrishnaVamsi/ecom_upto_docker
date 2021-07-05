//Models/UserModel.go
package Models

type Product struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `gorm:"type:int" json:"quantity"`
}

func (b *Product) TableName() string {
	return "product"
}
