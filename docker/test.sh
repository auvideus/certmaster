#!/bin/bash

echo -e "\n-----> BUILDING SERVER BASE\n"
docker build -t certmaster-server -f docker/Dockerfile_server .
echo -e "\n-----> BUILDING SERVER TEST\n"
docker build -t certmaster-server-test -f docker/test/Dockerfile_server .
echo -e "\n-----> RUNNING SERVER TEST\n"
docker run -t certmaster-server-test
echo -e "\n-----> BUILDING CLIENT BASE\n"
docker build -t certmaster-client -f docker/Dockerfile_client .
echo -e "\n-----> BUILDING CLIENT TEST\n"
docker build -t certmaster-client-test -f docker/test/Dockerfile_client .
echo -e "\n-----> RUNNING CLIENT TEST\n"
docker run -t certmaster-client-test
