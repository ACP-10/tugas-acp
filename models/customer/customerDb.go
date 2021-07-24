package customer

import "time"

type Customer struct {
	CustomerId int       `json:"customerId" gorm:"primaryKey"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}