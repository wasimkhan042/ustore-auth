#!/bin/bash

if [ "$(docker ps -q -f name=ustore-auth-mysql-db)" ]; then
  docker rm -f ustore-auth-mysql-db
fi

if [ "$(docker ps -q -f name=ustore-auth-mongo-db)" ]; then
  docker rm -f ustore-auth-mongo-db
fi