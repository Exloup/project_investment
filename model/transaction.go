package model

// Struct for trx_product
type TransactionProducts struct {
	Id                string `json:"id"`
	TrxCode           string `json:"trx_code"`
	DateTransaction   string `json:"date"`
	PortofolioId      string `json:"portofolio_id"`
	ProductId         string `json:"product_id"`
	StatusDescription string `json:"status_description"`
}

// Struct for trx_detail
type TransactionDetails struct {
	Id           string `json:"id"`
	PriceId      string `json:"price_id"`
	TrxProductId string `json:"trx_product_id"`
	Quantity     string `json:"quantity"`
}
