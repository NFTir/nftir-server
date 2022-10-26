include .env
SERVER_IMAGE=logann131/nftir-server:$(SERVER_TAG)
SERVER_TAG=1.0
CONTAINER_NAME=NFTir-server
DOCKER_RUN_SCRIPT = docker run -d --rm $\
		    		--name NFTir-server $\
		    		--env-file .env $\
		    		-e GIN_MODE=release $\
		    		-p $(PORT):$(PORT) $\
		    		$(SERVER_IMAGE)

DOCKER_RM_SCRIPT  = docker rm -f $(CONTAINER_NAME)
DOCKER_RMI_SCRIPT = docker rmi -f $(SERVER_IMAGE)

## login
docker-login:
	@echo docker logging in as logann131...
	@docker login -u logann131

## pull image
docker-pull:
	docker pull $(SERVER_IMAGE)

## clean app
docker-clean:
	$(DOCKER_RM_SCRIPT) $&
	$(DOCKER_RMI_SCRIPT)

## run image
docker-run:
	$(DOCKER_RUN_SCRIPT)

## logout
docker-logout:
	docker logout

## rm .env
remove-env:
	@echo removing .env...
	@rm .env

## build app -- ULTIMATE
.PHONY: build-app
build-app: docker-login docker-clean docker-pull docker-run docker-logout remove-env