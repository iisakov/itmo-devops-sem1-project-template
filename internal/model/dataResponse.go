package model

type DataResponse struct {
	TotalItems      int     `json:"total_items" db:"total_items" csv:"total_items"`
	TotalCategories int     `json:"total_categories" db:"total_categories" csv:"total_categories"`
	TotalPrice      float64 `json:"total_price" db:"total_price" csv:"total_price"`
}
