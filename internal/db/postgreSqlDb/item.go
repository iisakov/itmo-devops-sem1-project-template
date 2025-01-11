package postgreSqlDb

import (
	"fmt"
	"price/internal/model"
)

func (ps PostgreSqlDb) GetItems() ([]string, error) {
	res, err := ps.db.Query("select table_name from information_schema.tables")
	if err != nil {
		return []string{}, err
	}
	names := make([]string, 0)
	for res.Next() {
		var name string
		err = res.Scan(&name)
		if err != nil {
			return []string{}, err
		}
		fmt.Printf("tableName: %s \n", name)
		names = append(names, name)
	}

	return names, nil
}

func (ps PostgreSqlDb) AddItems(items model.Items) error {
	fmt.Println(items)
	return nil
}
