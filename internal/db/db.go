package db

import "price/internal/model"

type Storage interface {
	GetTotal() (model.DataResponse, error)
	AddItems(items model.Items) error
}
