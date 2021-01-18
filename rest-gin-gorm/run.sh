#!/bin/bash

echo "starting..."
PORT=8080 DB_URL="mybatis:mybatis@/rest_gin_gorm?charset=utf8&parseTime=True&loc=Local" ./rest-gin-gorm
