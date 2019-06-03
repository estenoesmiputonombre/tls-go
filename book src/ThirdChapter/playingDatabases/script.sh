#!/usr/bin/env bash

CONTAINER_DATABASE=mysqlgo
CONTAINER_GO=logicgo
IMAGE_DATABASE=mysql
IMAGE_GO=golang
TAG=latest
DATASTORE=/docker-entrypoint-initdb.d
NETWORK=mysqlgo

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then 
        docker stop $CONTAINER_DATABASE
    fi
    docker CONTAINER_DATABASE rm $CONTAINER_DATABASE     
fi

docker run --detach \
           --volume "${PWD}/database:${DATASTORE}" \
           --env-file "${PWD}/database/secret" \
           --name ${CONTAINER_DATABASE} \
           --publish 3000-4000:3306 \
           --network ${NETWORK} ${IMAGE_DATABASE}:${TAG}

if [ $(docker CONTAINER_DATABASE ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then
    if [ $(docker CONTAINER_DATABASE ls --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then 
        MYSQL_ROOT_ADDR=$(docker CONTAINER_DATABASE inspect mysqlgo --format "{{.NetworkSettings.Networks.golang_mysql.IPAddress}}")
        MYSQL_ROOT_PORT=$(docker CONTAINER_DATABASE inspect mysqlgo --format '{{ (index (index .NetworkSettings.Ports "3306/tcp") 0).HostPort }}')
        MYSQL_ROOT_PASSWORD=$(echo $(docker exec -it ${CONTAINER_DATABASE} bash -c 'echo "$MYSQL_ROOT_PASSWORD"'))
        if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_GO) -ge 1 ]; then
            if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_GO) -ge 1 ]; then 
                docker stop $CONTAINER_GO
            fi
            docker container rm $CONTAINER_GO 
        fi
        docker run --detach \
                   --env-file "${PWD}/secret" \
                   --name ${CONTAINER_GO} \
                   --publish 8080:8080 \
                   --network ${NETWORK} ${IMAGE_GO}:${TAG}
    fi
fi