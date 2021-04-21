.PHONY: all

VERSION := 0.0.1
IMG_NAME := ventilo/dummy-service

docker-build:
	@docker build . -t $(IMG_NAME)
	@docker tag $(IMG_NAME):$(VERSION) $(IMG_NAME):latest

docker-push:
	@docker push $(IMG_NAME):$(VERSION)
	@docker push $(IMG_NAME):latest
