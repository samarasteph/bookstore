#!/bin/zsh
docker build . --target=app-stage --build-arg "SERVER_PORT=5000" --build-arg "DB_SERVER=192.168.0.45" -t bookstore-server:0.0.0.1
