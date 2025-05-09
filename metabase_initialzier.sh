#!/bin/bash

sudo apt-get update
sudo apt-get upgrade -y

sudo apt-get install -y docker.io

sudo systemctl start docker
sudo systemctl enable docker

sudo mkdir -p /var/lib/metabase
sudo chmod -R 777 /var/lib/metabase

docker stop metabase || true
docker rm metabase || true

docker run -d \
  -p 3000:3000 \
  -v /var/lib/metabase:/metabase.db \
  -e MB_DB_FILE=/metabase.db/metabase.db \
  -e JAVA_TOOL_OPTIONS="-Xms512m -Xmx512m" \
  -e MB_DB_AUTOMATIC_UPDATES=false \
  -e MB_EMBEDDED_DATABASE_AUTO_BACKUP=false \
  -e MB_DISABLE_SAMPLE_DATA=true \
  -e MB_ANON_TRACKING_ENABLED=false \
  --name metabase \
  metabase/metabase
