build: build-sin build-triangle build-square

build-sin:
	sudo docker build --no-cache -t sin-wave --build-arg SHAPE=config-sin.yaml .

build-triangle:
	sudo docker build --no-cache -t triangle-wave --build-arg SHAPE=config-triangle.yaml .

build-square:
	sudo docker build --no-cache -t square-wave --build-arg SHAPE=config-square.yaml .

run-sin:
	sudo docker run -d --rm --name sin-wave sin-wave

run-triangle:
	sudo docker run -d --rm --name triangle-wave triangle-wave

run-square:	
	sudo docker run -d --rm --name square-wave square-wave

kill-all:
	sudo docker ps -a -q | xargs sudo docker kill	
