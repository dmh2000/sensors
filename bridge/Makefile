all: build docker-build

build: $(wildcard *.go)
	@echo "===================================== Building rabbitmq bridge"
	staticcheck .
	go build -o bridge

docker-build: 
	@echo "===================================== Building rabbitmq bridge docker image"
	staticcheck .
	sudo docker build --quiet --no-cache -t bridge .

docker-run: 
	sudo docker run -d --rm --name bridge bridge

docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill	

clean:
	@echo "===================================== Cleaning bridge"
	-rm bridge
