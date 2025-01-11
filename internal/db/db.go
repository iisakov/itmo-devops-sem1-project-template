package db

import "price/internal/model"

type Storage interface {
	GetItems() ([]string, error)
	AddItems(items model.Items) error
}
