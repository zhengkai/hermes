#!/bin/bash

TARGET="Freya"

if [ "$HOSTNAME" != "$TARGET" ]; then
	>&2 echo only run in server "$TARGET"
	exit 1
fi

sudo docker stop hermes
sudo docker rm hermes
sudo docker rmi hermes

sudo cat /tmp/docker-hermes.tar | sudo docker load

sudo docker run -d --name hermes \
	--env "TANK_MYSQL=hermes:hermes@tcp(172.17.0.1:3306)/hermes" \
	--env "STATIC_DIR=/tmp" \
	--env "OUTPUT_PATH=/output" \
	--mount type=bind,source=/www/hermes/output,target=/output \
	--mount type=bind,source=/www/hermes/log,target=/log \
	--mount type=bind,source=/www/hermes/static,target=/tmp \
	--restart always \
	hermes
