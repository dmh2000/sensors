all: build docker-build

build: $(wildcard *.go)
	@echo "===================================== Building api"
	staticcheck .
	go build -o api

docker-build: 
	@echo "===================================== Building api docker image"
	sudo docker build --quiet --no-cache -t api .

docker-run: 
	sudo docker run -d --rm --name api api

docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill	

clean:
	@echo "===================================== Cleaning api"
	-rm api
