package domain

type SalesData struct {
	Model
	Day        string  `json:"day"`
	TotalPrice int     `json:"total_price"`
	Up         bool    `json:"up"`
	Rate       float64 `json:"rate"`
}
