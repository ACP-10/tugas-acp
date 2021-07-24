package product

import "time"

type Product struct {
	ProductId    int       `json:"productId" gorm:"primaryKey"`
	ProductName  string    `json:"productName"`
	ProductPrice int       `json:"productPrice"`
	Stock        int       `json:"stock"`
	CategoryId   int       `json:"categoryId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}