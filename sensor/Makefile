all: sensor docker-build

sensor: $(wildcard *.go)
	@echo "===================================== Building sensor"
	staticcheck .
	go build -o sensor

docker-build: docker-sin docker-triangle docker-square

docker-run: run-sin run-triangle run-square

docker-kill:
	@echo "===================================== Killing all containers"
	sudo docker ps -a -q | xargs sudo docker kill	

docker-sin:
	@echo "===================================== Building sin-wave docker image"
	sudo docker build --quiet --no-cache -t sin-wave --build-arg SHAPE=config-sin.yaml .

docker-triangle:
	@echo "===================================== Building triangle-wave docker image"
	sudo docker build --quiet --no-cache -t triangle-wave --build-arg SHAPE=config-triangle.yaml .

docker-square:
	@echo "===================================== Building square-wave docker image"
	sudo docker build --quiet --no-cache -t square-wave --build-arg SHAPE=config-square.yaml .

run-sin:
	sudo docker run -d --rm --name sin-wave sin-wave

run-triangle:
	sudo docker run -d --rm --name triangle-wave triangle-wave

run-square:	
	sudo docker run -d --rm --name square-wave square-wave

clean:
	@echo "=====================================  Cleaning sensor"
	-rm sensor