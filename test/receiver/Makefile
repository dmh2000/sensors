all: receiver  docker-build

receiver: $(wildcard *.go)
	staticcheck .
	go build -o receiver

docker-build: 
	sudo docker build --no-cache -t receiver .

docker-run:
	sudo docker run -it --rm --name receiver receiver


docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill	

clean:
	rm receiver