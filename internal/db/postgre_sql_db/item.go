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

	// Получаем данные из базы
	res, err := ps.db.Query(db.GetItems)
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

func (ps PostgreSqlDb) AddItems(items model.Items) (model.DataResponse, error) {
	response := model.DataResponse{}

	tx, err := ps.db.Begin()
	if err != nil {
		return response, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	for _, item := range items {
		_, err = tx.Exec(db.AddItem,
			item.CreateDate,
			item.Name,
			item.Category,
			item.Price)
		if err != nil {
			return response, err
		}
		response.TotalItems += 1
	}

	// Получаем данные из базы
	res := tx.QueryRow(db.TotalQuery)

	// Записываем ответ
	err = res.Scan(
		&response.TotalCategories,
		&response.TotalPrice)
	if err != nil {
		return response, err
	}

	err = tx.Commit()
	if err != nil {
		return response, err
	}

	return response, nil
}
