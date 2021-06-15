# Runs the standard make commands inside src/ but inside docker
# add DOCKER_BUILDKIT=0 to see problems
.PHONY: all
 
all: build test docs run

build:
	DOCKER_BUILDKIT=0 docker build . -t bcg-api --target app

test: build
	DOCKER_BUILDKIT=0 docker run bcg-api "make test"

docs:
	DOCKER_BUILDKIT=0 docker build --target docs --output type=local,dest=src/OUTTT -t bcg-docs .

run: build
	DOCKER_BUILDKIT=0 docker run -p 3000:3000 bcg-app
