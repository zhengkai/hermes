SHELL:=/bin/bash

-include ./server/build/config.ini

play:
	./build/build-server.sh $(type)
	/www/hermes/dist/dev/hermes-server-next -s 80x50 /www/hermes/static/forza5.mp4 
	# /www/hermes/server/dist/dev/hermes-server-next /www/hermes/static/1489962262890565646.mp4

start:
	# ./server/build/build-server.sh $(type)
	# /www/hermes/server/dist/dev/hermes-server-next
	./build/run-server.sh $(type)

stop:
	./build/stop-server.sh $(type)

dev:
	./build/run-server.sh dev

stopdev:
	./build/stop-server.sh dev

prod:
	./build/run-server.sh prod

stopprod:
	./build/stop-server.sh prod
