#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u admin -p admin workManager < "/docker-entrypoint-initdb.d/001-create-table_types.sql"