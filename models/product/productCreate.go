package product

type ProductCreate struct {
	ProductName  string    `json:"productName"`
	ProductPrice int       `json:"productPrice"`
	Stock        int       `json:"stock"`
	CategoryId   int       `json:"categoryId"`
}