
docker-build: 
	@echo "===================================== Building mosquitto docker image"
	sudo docker build --quiet -t mosquitto .

docker-run:
	sudo docker run -it --rm --name mosquitto -p 1883:1883 mosquitto
