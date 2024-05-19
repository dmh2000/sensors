all:
	@echo "===================================== Building rabbitmq docker image"
	@make docker-build
	@echo "===================================== Running rabbitmq docker image"
	@make docker-run

docker-build: 
	@echo "===================================== Building rabbitmq docker image"
	sudo docker build --quiet -t rabbitmq .

docker-run:
	sudo docker run -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management

