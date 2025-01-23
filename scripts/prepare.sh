#!/bin/bash

# Для прерывания скрипта в случае возникновения ошибок
set -e

echo "Читаем что понаписанов .env ..."
# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs)

echo "Создаём таблицу price..."
PGPASSWORD=$USER_PASSWORD psql -U $USER_NAME -h $DB_HOST_OUTSIDE -p $DB_PORT -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS prices (
    id SERIAL PRIMARY KEY,
    create_date DATE NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
create or replace function add_items_and_calc_stat(
 	in create_date date, in name varchar, in category varchar, in price numeric, out total_items integer, out total_categories integer, out total_price numeric) as  \$\$
 	begin
 		insert into prices (
 			\"create_date\",
 			\"name\",
 			\"category\",
 			\"price\")
 		values (
 			add_items_and_calc_stat.create_date,
 			add_items_and_calc_stat.name,
 			add_items_and_calc_stat.category,
 			add_items_and_calc_stat.price);

 		select
 			count(*),
 			count(distinct p.category),
 			sum(p.price)
 		into total_items, total_categories, total_price
 		from prices p;
 	end
 \$\$ language plpgsql;
 "

echo "Готовы к работе!"