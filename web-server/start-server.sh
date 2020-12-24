#!/bin/sh
IP_ADDR=$(ip -o -4  addr | grep $NIC_NAME_ENV | sed -r 's/^.+inet ([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+).+$/\1/')
/app/webserver --port ${SERVER_PORT_ENV} --addr $IP_ADDR --dbserver ${DB_SERVER_ENV}
