#!/bin/bash

export DOCKER_USERNAME=auvideus
echo -e "\n-----> BUILDING SERVER BASE\n"
docker build -t $DOCKER_USERNAME/certmaster-server -f docker/Dockerfile_server .
echo -e "\n-----> BUILDING SERVER TEST\n"
docker build -t certmaster-server-test -f docker/test/Dockerfile_server .
echo -e "\n-----> RUNNING SERVER TEST\n"
docker run -t certmaster-server-test
echo -e "\n-----> BUILDING CLIENT BASE\n"
docker build -t $DOCKER_USERNAME/certmaster-client -f docker/Dockerfile_client .
echo -e "\n-----> BUILDING CLIENT TEST\n"
docker build -t certmaster-client-test -f docker/test/Dockerfile_client .
echo -e "\n-----> RUNNING CLIENT TEST\n"
docker run -t certmaster-client-test

if [[ -v DOCKER_PASSWORD ]]; then
    docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD"
    docker push "$DOCKER_USERNAME"/certmaster-server
    docker push "$DOCKER_USERNAME"/certmaster-client
fi
