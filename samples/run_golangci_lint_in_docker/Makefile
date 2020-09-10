SHELL:=/bin/bash

.PHONY: all
all:

.PHONY: up
up:
	docker-compose up --force-recreate --always-recreate-deps --renew-anon-volumes --remove-orphans

.PHONY: down
down:
	docker-compose down --volumes --remove-orphans
