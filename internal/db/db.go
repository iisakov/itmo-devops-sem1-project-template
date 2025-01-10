package db

import "price/internal/model"

type Storage interface {
	GetItems() (model.Items, error)
	AddItems(items model.Items) error
}
