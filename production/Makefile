include .env
APP=server
SERVER_IMAGE=logann131/nftir-server:$(SERVER_TAG)
SERVER_TAG=1.0
DOCKER_RMI=docker rmi -f
DOCKER_PUSH=docker push
DOCKER_COMPOSE_BUILD=docker compose build --no-cache
DOCKER_COMPOSE_DOWN=docker compose down
DOCKER_COMPOSE_UP=docker compose up -d

## Production mode
.PHONY: start-image
start-image: remove-image build-image

.PHONY: remove-image
remove-image:
	$(DOCKER_RMI) $(SERVER_IMAGE)

.PHONY: build-image
build-image:
	$(DOCKER_COMPOSE_BUILD)


.PHONY: start-container
start-container: down-container up-container


.PHONY: down-container
down-container:
	$(DOCKER_COMPOSE_DOWN)

.PHONY: up-container
up-container:
	$(DOCKER_COMPOSE_UP)


## docker login
docker-login:
	@echo docker logging in as logann131...
	@docker login -u logann131

## docker logout
docker-logout:
	docker logout

## Clean app
.PHONY: clean-app
clean-app: down-container remove-image

## build app -- ULTIMATE
.PHONY: build-app
build-app: docker-logout docker-login start-image start-container docker-logout