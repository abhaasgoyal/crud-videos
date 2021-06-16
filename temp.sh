#!/usr/bin/env bash
set -euo pipefail

# Create and run database instance as root
docker run --add-host=host.docker.internal:host-gateway --name=mysql1 -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 -d mysql
# Show database
docker exec -it mysql1 mysql -uroot -p

# create
CREATE USER 'abhaas'@'localhost' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON * . * TO 'abhaas'@'localhost';
CREATE DATABASE video_database;
