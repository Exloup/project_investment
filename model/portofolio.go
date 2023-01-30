package model

// Sturct for Portofolio
type Portofolios struct {
	Id            string `json:"id"`
	AccountId     string `json:"account_id"`
	TotalQuantity string `json:"total_quantity"`
	Profit        string `json:"profit"`
}
