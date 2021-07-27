#!/bin/zsh

docker run --rm --name dbserver --hostname dbserver bookstore/db:latest
#docker run --detach --rm --name dbserver --hostname dbserver bookstore/db:latest
