path = "here"

# build docker base image
setup:
	./scripts/build-base-image.sh
.PHONY: setup

# build all services
build:
	docker-compose up -d --build ${#@D}
.PHONY: build

all: 
	

clean: