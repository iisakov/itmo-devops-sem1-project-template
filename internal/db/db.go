package db

import "price/internal/model"

type Storage interface {
	GetTotal() (model.DataResponse, error)
	GetItems() (model.Items, error)
	AddItems(items model.Items) (model.DataResponse, error)
}
