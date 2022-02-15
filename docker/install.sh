#!/bin/bash

TARGET="Jack"

if [ "$HOSTNAME" != "$TARGET" ]; then
	>&2 echo only run in server "$TARGET"
	exit 1
fi

sudo docker stop hermes
sudo docker rm hermes
sudo docker rmi hermes

sudo cat /tmp/docker-hermes.tar | sudo docker load

sudo docker run -d --name hermes \
	--mount type=bind,source=/www/hermes/static,target=/static \
	--restart always \
	-p 0.0.0.0:23:30023 \
	hermes
