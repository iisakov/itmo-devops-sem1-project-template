package postgreSqlDb

import (
	"price/internal/db"
	"price/internal/model"
)

func (ps PostgreSqlDb) GetTotal() (model.DataResponse, error) {
	response := model.DataResponse{}

	tx, err := ps.db.Begin()
	if err != nil {
		return response, err
	}
	defer func() { _ = tx.Rollback() }()

	// Получаем данные из базы
	res := tx.QueryRow(db.TotalQuery)

	// Записываем ответ
	err = res.Scan(
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

	// Открываем транзакцию
	tx, err := ps.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }() // Научился красиво обрабатывать параметры в defer

	// Получаем данные из базы
	res, err := tx.Query(db.GetItems)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Close() }()

	for res.Next() {
		var item model.Item
		err = res.Scan(
			&item.Id,
			&item.CreateDate,
			&item.Name,
			&item.Category,
			&item.Price)
		if err != nil {
			return nil, err
		}
		response = append(response, item)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}

	// Записываем ответ
	return response, nil
}

func (ps PostgreSqlDb) AddItems(items model.Items) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, item := range items {
		_, err := tx.Exec(db.AddItem,
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
