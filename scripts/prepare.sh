#!/bin/bash

# Для прерывания скрипта в случае возникновения ошибок
set -e

echo "Читаем что понаписанов .env ..."
export $(grep -v '^#' .env | xargs)

echo "Создаём таблицу price..."
PGPASSWORD=$USER_PASSWORD psql -U $USER_NAME -h $DB_HOST_OUTSIDE -p $DB_PORT -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS prices (
    id SERIAL PRIMARY KEY,
    manufacture_date DATE NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_category VARCHAR(255) NOT NULL,
    product_price DECIMAL(10,2) NOT NULL
);"

echo "Готовы к работе!"