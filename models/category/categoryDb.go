package category

import "time"

type Category struct {
	CategoryId   int       `json:"categoryId" gorm:"primaryKey"`
	CategoryName string    `json:"categoryName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
