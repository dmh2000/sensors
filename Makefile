all: build

build:
	$(MAKE) -C sensor
	$(MAKE) -C bridge
	$(MAKE) -C api
	$(MAKE) -C dashboard

docker-build:
	$(MAKE) -C sensor docker-build
	$(MAKE) -C bridge docker-build
	$(MAKE) -C api docker-build
	$(MAKE) -C dashboard docker-build

docker-run:
	$(MAKE) -C rabbitmq rabbit
	$(MAKE) -C sensor docker-run
	$(MAKE) -C bridge docker-run
	$(MAKE) -C api docker-run
	$(MAKE) -C dashboard docker-run
	sudo docker ps

docker-compose:
	# build the images
	sudo docker compose build
	# start the services
	sudo docker compose up
	# start the sensors outside of compose
	$(MAKE) -C sensor docker-build
	$(MAKE) -C sensor docker-run


docker-kill:
	-sudo docker ps -a -q | xargs sudo docker kill

docker-clean:
	-sudo docker ps -a -q | xargs sudo docker kill
	-sudo docker  container list -aq | xargs sudo docker container rm
	-sudo docker  image list -q | xargs sudo docker image rm

clean:
	$(MAKE) -C sensor clean
	$(MAKE) -C api clean
	$(MAKE) -C dashboard clean
