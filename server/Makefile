include .env
## CompileDaemon for hot reload
DEV_MODE_PATH=~/go/bin/CompileDaemon ##https://github.com/githubnemo/CompileDaemon
DEV_MODE_CMD=-command
APP=server
SERVER_IMAGE=logann131/nftir-server:$(SERVER_TAG)
SERVER_TAG=1.0
DOCKER_RMI=docker rmi -f
DOCKER_RM=docker rm -f
DOCKER_PUSH=docker push
DOCKER_IMAGE_LIST_ID=docker images -q
DOCKER_CONTAINER_LIST_ID=docker ps -aq
DOCKER_BUILD_SCRIPT = 	docker build $\
						--build-arg PORT=$(PORT) $\
						--no-cache $\
						-t $(SERVER_IMAGE) .
DOCKER_RUN_SCRIPT = docker run -d --rm $\
					--env-file .env $\
					-e GIN_MODE=release $\
					-p $(PORT):$(PORT) $\
					$(SERVER_IMAGE)
DOCKER_COMPOSE_BUILD=docker compose build --no-cache
DOCKER_COMPOSE_DOWN=docker compose down
DOCKER_COMPOSE_UP=docker compose up -d


 ####### LOCAL ######
.PHONY: go-build-local
go-build-local: clean server run

.PHONY: clean
clean: 
	go $@

build: server
	go $@ -o $< .

run: server
	./$<

.PHONY: dev-mode
dev-mode:
	$(DEV_MODE_PATH) $(DEV_MODE_CMD)="./$(APP)"

 
####### DOCKER ######
.PHONY: docker-build
docker-build: remove-image
	$(DOCKER_BUILD_SCRIPT)

docker-run: 
	$(DOCKER_RUN_SCRIPT)

docker-clean:
	$(DOCKER_RM) $$($(DOCKER_CONTAINER_LIST_ID)) $&
	$(DOCKER_RMI) $$($(DOCKER_IMAGE_LIST_ID)) 

## -------- ULTIMATE -----------
docker-server: docker-build docker-run


## Push app to the hub
.PHONY: push-image
push-image: start-image 
	$(DOCKER_PUSH) $(SERVER_IMAGE)


########## DOCKER COMPOSE ########
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

## Clean app
.PHONY: clean-app
clean-app: down-container remove-image

## build app
.PHONY: build-app
build-app: start-image start-container