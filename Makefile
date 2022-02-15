SHELL:=/bin/bash

-include ./server/build/config.ini

play:
	./build/build-server.sh $(type)
	mv ./dist/dev/hermes-next ./dist/dev/hermes
	./dist/dev/hermes -version
	# ./dist/dev/hermes -size 20x20 -verbose ./static/forza5.mp4
	# ./dist/dev/hermes -size 16x16 -verbose ./static/1489962262890565646.mp4

release:
	./build/pack.sh linux amd64
	./build/pack.sh linux arm64
	./build/pack.sh darwin amd64
	./build/pack.sh darwin arm64

install:
	./build/build-server.sh prod
	sudo cp ./dist/prod/hermes /usr/local/bin/

start:
	./build/run-server.sh $(type)

stop:
	./build/stop-server.sh $(type)
