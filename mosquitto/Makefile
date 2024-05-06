all:
	@echo "===================================== Building mosquitto docker image"
	@make docker-build
	@echo "===================================== Running mosquitto docker image"
	@make docker-run

docker-build: 
	@echo "===================================== Building mosquitto docker image"
	sudo docker build --quiet -t mosquitto .

docker-run:
	sudo docker run -d --rm --name mosquitto -p 1883:1883 mosquitto
