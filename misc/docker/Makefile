SHELL:=/bin/bash

build: git
	sudo docker build -t hermes -f Dockerfile ../..

run: build
	sudo docker run --env "HERMES_MYSQL=hermes:hermes@tcp(172.17.0.1:3306)/hermes" \
		--mount type=bind,source=/www/hermes/tmp,target=/tmp \
		--mount type=bind,source=/www/hermes/server/dist/prod/log,target=/log \
		hermes

install: build
	sudo docker save hermes > docker-hermes.tar
	scp docker-hermes.tar freya:/tmp
	scp install.sh freya:/tmp
	ssh freya "chmod +x /tmp/install.sh && /tmp/install.sh && rm /tmp/install.sh"

git:
	../../server/build/git-hash.sh > ../../server/build/.git-hash

clean:
	sudo docker stop hermes || :
	sudo docker rm hermes || :
	sudo docker rmi hermes || :
	sudo docker image prune --force
