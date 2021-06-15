#!/usr/bin/env bash
cd ../..
migrate -source file://db/sql -database 'mysql://root:root@tcp(127.0.0.1:3306)/workManager' up