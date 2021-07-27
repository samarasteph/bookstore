#!/bin/zsh

docker image prune -f
docker build . --target web-stage --build-arg "SERVER_PORT=5000" --build-arg "DB_SERVER=dbserver" -t  bookstore/web:0.0.0.1
