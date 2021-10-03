package model

// RazorPayResponse .
type RazorPayResponse struct {
	Count int            `json:"count"`
	Items []RazorPayItem `json:"items"`
}

// RazorPayItem .
type RazorPayItem struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
