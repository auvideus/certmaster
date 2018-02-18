#!/bin/bash

docker build -t certmaster-server -f docker/Dockerfile_server
docker build -t certmaster-server-test -f docker/test/Dockerfile_server
docker run -t certmaster-server
docker build -t certmaster-client -f docker/Dockerfile_client
docker build -t certmastesr-client-test -f docker/test/Dockerfile_client
docker run -t certmaster-client
