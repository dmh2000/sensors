build: sin triangle square

sin:
	sudo docker build --no-cache -t sin-wave --build-arg SHAPE=config-sin.yaml .

triangle:
	sudo docker build --no-cache -t triangle-wave --build-arg SHAPE=config-triangle.yaml .

square:
	sudo docker build --no-cache -t square-wave --build-arg SHAPE=config-square.yaml .

run:
	sudo docker run -d --rm --name sin-wave sin-wave
	sudo docker run -d --rm --name triangle-wave triangle-wave
	sudo docker run -d --rm --name square-wave square-wave

