sensors: sin-wave triangle-wave square-wave

sin-wave: FORCE
	@echo build sin-wave
	staticcheck .
	cp config-sin.yaml config.yaml
	go build -o sin-wave
	./sin-wave &

triangle-wave: FORCE
	@echo build triangle-wave
	staticcheck .
	cp config-triangle.yaml config.yaml
	go build -o triangle-wave
	./triangle-wave &

square-wave: FORCE
	@echo build square-wave
	staticcheck .
	cp config-square.yaml config.yaml
	go build -o square-wave
	./square-wave &

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
	-rm sin-wave square-wave triangle-wave

FORCE:
