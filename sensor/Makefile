all: sensor docker-build

sensor: $(wildcard *.go)
	staticcheck .
	go build -o sensor

docker-build: docker-sin docker-triangle docker-square

docker-run: run-sin run-triangle run-square

docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill	

docker-sin:
	sudo docker build --no-cache -t sin-wave --build-arg SHAPE=config-sin.yaml .

docker-triangle:
	sudo docker build --no-cache -t triangle-wave --build-arg SHAPE=config-triangle.yaml .

docker-square:
	sudo docker build --no-cache -t square-wave --build-arg SHAPE=config-square.yaml .

run-sin:
	sudo docker run -d --rm --name sin-wave sin-wave

run-triangle:
	sudo docker run -d --rm --name triangle-wave triangle-wave

run-square:	
	sudo docker run -d --rm --name square-wave square-wave

clean:
	rm sensor