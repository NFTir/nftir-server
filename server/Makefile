## CompileDaemon for hot reload
DEV_MODE_PATH=~/go/bin/CompileDaemon ##https://github.com/githubnemo/CompileDaemon
DEV_MODE_CMD=-command
APP=server
SERVER_IMAGE=nnguyen6/nftir-server:$(SERVER_TAG)
SERVER_TAG=1.0
DOCKER_RMI=docker rmi -f
DOCKER_COMPOSE_BUILD=docker compose build --no-cache
DOCKER_COMPOSE_DOWN=docker compose down
DOCKER_COMPOSE_REMOVE=docker compose
DOCKER_COMPOSE_UP=docker compose up -d


## Development mode
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


## Clean app
.PHONY: clean-app
clean-app: down-container remove-image

## build app
.PHONY: build-app
build-app: start-image start-container