all:
	$(MAKE) -c receiver
	$(MAKE) -c sender

docker-run:
	$(MAKE) -c receiver docker-run
	$(MAKE) -c sender docker-run


docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill
