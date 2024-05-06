sensors: sin-wave triangle-wave square-wave

sin-wave: FORCE
	@echo build sin-wave
	staticcheck .
	cp config-sin.yaml config.yaml
	go build -o sin-wave

triangle-wave: FORCE
	@echo build triangle-wave
	staticcheck .
	cp config-triangle.yaml config.yaml
	go build -o triangle-wave

square-wave: FORCE
	@echo build square-wave
	staticcheck .
	cp config-square.yaml config.yaml
	go build -o square-wave

run: run-sin run-triangle run-square

run-sin:
	./sin-wave config-sin.yaml &

run-triangle:
	./triangle-wave config-triangle.yaml &

run-square:	
	./square-wave config-square.yaml &	

docker-build: docker-build-sin docker-build-triangle docker-build-square

docker-run: docker-run-sin docker-run-triangle docker-run-square

docker-kill:
	@echo "===================================== Killing all containers"
	sudo docker ps -a -q | xargs sudo docker kill	

docker-build-sin:
	@echo "===================================== Building sin-wave docker image"
	staticcheck .
	sudo docker build --quiet --no-cache -t sin-wave --build-arg SHAPE=config-sin.yaml .

docker-build-triangle:
	@echo "===================================== Building triangle-wave docker image"
	staticcheck .
	sudo docker build --quiet --no-cache -t triangle-wave --build-arg SHAPE=config-triangle.yaml .

docker-build-square:
	@echo "===================================== Building square-wave docker image"
	staticcheck .
	sudo docker build --quiet --no-cache -t square-wave --build-arg SHAPE=config-square.yaml .

docker-run-sin:
	sudo docker run -d --rm --name sin-wave sin-wave

docker-run-triangle:
	sudo docker run -d --rm --name triangle-wave triangle-wave

docker-run-square:	
	sudo docker run -d --rm --name square-wave square-wave

clean:
	@echo "=====================================  Cleaning sensor"
	-rm sin-wave square-wave triangle-wave

FORCE:
