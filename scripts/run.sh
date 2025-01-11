#!/bin/bash

# Для прерывания скрипта в случае возникновения ошибок
set -e

docker compose down
docker image rm itmo-devops-sem1-project-template-backend -f
docker compose up -d
