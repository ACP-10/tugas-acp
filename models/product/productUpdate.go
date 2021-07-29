package product

type ProductUpdate struct {
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Stock        int    `json:"stock"`
	CategoryId   int    `json:"categoryId"`
}
