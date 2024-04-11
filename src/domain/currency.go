package domain

type Currency struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
	ADate string  `json:"aDate"`
}
