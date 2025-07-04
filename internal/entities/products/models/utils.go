package models

type ProductFilter struct {
	ProductIDs []string `json:"productIDs"`
	Category   string   `json:"category"`
}
