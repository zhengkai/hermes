SHELL:=/bin/bash

-include ./server/build/config.ini

play:
	./build/build-server.sh $(type)
	mv /www/hermes/dist/dev/hermes-next /www/hermes/dist/dev/hermes
	# /www/hermes/dist/dev/hermes-server -s 20x20 -v /www/hermes/static/forza5.mp4
	/www/hermes/dist/dev/hermes -size 16x16 -verbose /www/hermes/static/1489962262890565646.mp4

release:
	GOOS=linux GOARCH=arm64 ./build/build-server.sh release
	./build/pack.sh linux_arm64
	GOOS=linux GOARCH=amd64 ./build/build-server.sh release
	./build/pack.sh linux_amd64
	GOOS=darwin GOARCH=arm64 ./build/build-server.sh release
	./build/pack.sh mac_arm64
	GOOS=darwin GOARCH=amd64 ./build/build-server.sh release
	./build/pack.sh mac_amd64

prod:
	./build/build-server.sh prod
	sudo cp ./dist/prod/hermes /usr/local/bin/

start:
	./build/run-server.sh $(type)

stop:
	./build/stop-server.sh $(type)
