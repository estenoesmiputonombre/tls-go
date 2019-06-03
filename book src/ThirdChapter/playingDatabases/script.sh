#!/usr/bin/env bash

CONTAINER=mysqlgo
IMAGE=mysql
TAG=latest
DATASTORE=/docker-entrypoint-initdb.d

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then 
        docker stop $CONTAINER
    fi
    docker container rm $CONTAINER     
fi

docker run --detach \
           --volume "${PWD}/database:${DATASTORE}" \
           -e MYSQL_ROOT_PASSWORD=abc123. \
           --name ${CONTAINER} \
           --publish 3000-4000:3306 \
           --network golang_mysql ${IMAGE}:${TAG}