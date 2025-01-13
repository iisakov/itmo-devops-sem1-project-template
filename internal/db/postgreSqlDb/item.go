package postgreSqlDb

import (
	"fmt"
	"price/internal/db"
	"price/internal/model"
)

func (ps PostgreSqlDb) GetTotal() (model.DataResponse, error) {
	response := model.DataResponse{}

	// Получаем данные из базы
	res := ps.db.QueryRow(db.TotalQuery)

	// Записываем ответ
	err := res.Scan(
		&response.TotalItems,
		&response.TotalCategories,
		&response.TotalPrice)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ps PostgreSqlDb) AddItems(items model.Items) error {
	fmt.Println(items)
	return nil
}
