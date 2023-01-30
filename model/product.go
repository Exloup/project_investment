package model

// Struct for Product
type Products struct {
	Id                 int    `json:"id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"description_product"`
}

// Struct for Price
type Prices struct {
	Id        int    `json:"id"`
	Price     string `json:"price"`
	Date      string `json:"date"`
	ProductId string `json:"product_id"`
	Performa  string `json:"performa"`
}
