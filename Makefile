SHELL:=/bin/bash

-include ./server/build/config.ini

start:
	./server/build/build-server.sh $(type)
	/www/hermes/server/dist/dev/hermes-server-next
	# ./server/build/run-server.sh $(type)

stop:
	./server/build/stop-server.sh $(type)

dev:
	./server/build/run-server.sh dev

stopdev:
	./server/build/stop-server.sh dev

prod:
	./server/build/run-server.sh prod

stopprod:
	./server/build/stop-server.sh prod
