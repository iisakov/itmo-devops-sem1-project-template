package model

type Item struct {
	Id         int    `json:"id" db:"id" csv:"id"`
	Name       string `json:"name" db:"name"  csv:"name"`
	Category   string `json:"category" db:"category" csv:"category"`
	Price      int    `json:"price" db:"price" csv:"price"`
	CreateDate string `json:"create_date" db:"create_date" csv:"create_date"`
}

type Items []Item
