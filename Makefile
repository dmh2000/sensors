all:
	$(MAKE) -C receiver
	$(MAKE) -C sensor

docker-run:
	$(MAKE) -C receiver docker-run
	$(MAKE) -C sensor docker-run


docker-kill:
	sudo docker ps -a -q | xargs sudo docker kill

clean:
	$(MAKE) -C receiver clean
	$(MAKE) -C sensor clean
