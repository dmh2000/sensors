all: build docker-build

build: $(wildcard *.go)
	@echo "===================================== Building dashboard"
	staticcheck .
	go build -o dashboard

docker-build: 
	@echo "===================================== Building dashboard docker image"
	sudo docker build --quiet --no-cache -t dashboard .

docker-run: 
	sudo docker run -d --rm --name dashboard dashboard

docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill	

clean:
	@echo "===================================== Cleaning dashboard"
	-rm dashboard
