package postgreSqlDb

import (
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

func (ps PostgreSqlDb) GetItems() (model.Items, error) {
	response := model.Items{}

	// Получаем данные из базы
	res, err := ps.db.Query(db.GetItems)
	if err != nil {
		return response, err
	}
	for res.Next() {
		var item model.Item
		err = res.Scan(
			&item.Id,
			&item.CreateDate,
			&item.Name,
			&item.Category,
			&item.Price)
		if err != nil {
			return response, err
		}
		response = append(response, item)
	}
	// Записываем ответ
	return response, nil
}

func (ps PostgreSqlDb) AddItems(items model.Items) error {
	for _, item := range items {
		_, err := ps.db.Exec(db.AddItem,
			item.CreateDate,
			item.Name,
			item.Category,
			item.Price)
		if err != nil {
			return err
		}
	}
	return nil
}
