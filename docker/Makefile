SHELL:=/bin/bash

build: git
	sudo docker build -t hermes -f Dockerfile ..

run: build
	sudo docker run \
		--mount type=bind,source=/www/hermes/static,target=/static \
		-p 0.0.0.0:23:30023 \
		hermes

install: build
	sudo docker save hermes > docker-hermes.tar
	scp docker-hermes.tar jack:/tmp
	scp install.sh jack:/tmp
	ssh jack "chmod +x /tmp/install.sh && /tmp/install.sh && rm /tmp/install.sh"

git:
	../build/git-hash.sh > ../build/.git-hash

clean:
	sudo docker stop hermes || :
	sudo docker rm hermes || :
	sudo docker rmi hermes || :
	sudo docker image prune --force
