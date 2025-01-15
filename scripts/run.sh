#!/bin/bash

# Для прерывания скрипта в случае возникновения ошибок
set -e

go build -o ./bin/app ./cmd/price/price.go
./bin/app

