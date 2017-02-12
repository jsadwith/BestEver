NAME := best-ever
VERSION := $(shell git rev-parse --verify HEAD)

deps:
	godep get
	godep save

run:
	docker-compose build
	docker-compose up -d

shell:
	docker-compose run $(NAME) bash

down:
	docker-compose down

logs:
	docker-compose logs -f $(NAME)
