test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build house

stop:
	sudo docker-compose stop