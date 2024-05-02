all: build

build:
	$(MAKE) -C sensor
	$(MAKE) -C bridge
	$(MAKE) -C api
	$(MAKE) -C dashboard

docker-run:
	$(MAKE) -C rabbitmq rabbit
	$(MAKE) -C sensor docker-run
	$(MAKE) -C bridge docker-run
	$(MAKE) -C api docker-run
	$(MAKE) -C dashboard docker-run
	sudo docker ps

docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill

clean:
	$(MAKE) -C sensor clean
	$(MAKE) -C api clean
	$(MAKE) -C dashboard clean
