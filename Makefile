SHELL:=/bin/bash

-include ./server/build/config.ini

play:
	./build/build-server.sh $(type)
	mv ./dist/dev/hermes-next ./dist/dev/hermes
	./dist/dev/hermes || :
	# ./dist/dev/hermes -version
	# ./dist/dev/hermes -verbose -size 2x2 ./static/screenshot.webp
	# ./dist/dev/hermes -size 2x2 ./static/screenshot.webp
	./dist/dev/hermes -size 10x10 -seek 80 -verbose -frames 100 ./static/forza5.mp4
	# ./dist/dev/hermes -size 16x16 -verbose ./static/1489962262890565646.mp4

release:
	./build/pack.sh linux amd64
	./build/pack.sh linux arm64
	./build/pack.sh darwin amd64
	./build/pack.sh darwin arm64

install:
	./build/build-server.sh prod
	mv ./dist/prod/hermes-next ./dist/prod/hermes
	sudo cp ./dist/prod/hermes /usr/local/bin/

start:
	./build/run-server.sh $(type)

stop:
	./build/stop-server.sh $(type)
