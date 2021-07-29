package customer

import (
	"time"
)

type CustomerResponse struct {
	CustomerId int       `json:"customerId"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
