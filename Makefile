all: 

# run staticcheck and build
build:
	$(MAKE) -C sensor
	$(MAKE) -C bridge
	$(MAKE) -C api
	$(MAKE) -C dashboard

# build the docker containers
docker-build:
	$(MAKE) -C rabbitmq docker-build
	$(MAKE) -C sensor docker-build
	$(MAKE) -C mosquitto docker-build
	$(MAKE) -C bridge docker-build
	$(MAKE) -C api docker-build
	$(MAKE) -C dashboard docker-build

# launch the containers
docker-compose:
	# build the images
	sudo docker compose build
	# start the services
	sudo docker compose up

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
