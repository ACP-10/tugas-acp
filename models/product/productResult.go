package product

type ProductResult struct {
	ProductId    int    `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Stock        int    `json:"stock"`
	Category     string `json:"category"`
}
