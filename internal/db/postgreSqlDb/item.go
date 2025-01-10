package postgreSqlDb

import (
	"fmt"
	"price/internal/model"
)

func (ps PostgreSqlDb) GetItems() (model.Items, error) {
	return model.Items{}, nil
}

func (ps PostgreSqlDb) AddItems(items model.Items) error {
	fmt.Println(items)
	return nil
}
