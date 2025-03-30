package internal

type RateRequest struct {
	Service       string `json:"service"`
	RatePerMinute int    `json:"rpm"`
}