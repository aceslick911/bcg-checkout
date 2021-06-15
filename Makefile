# Runs the standard make commands inside src/ but inside docker
# add DOCKER_BUILDKIT=0 to see problems
.PHONY: all
 
all: build test docs run

build: # DOCKER_BUILDKIT=0 
	docker build --target build -t bcg-api . 

test: build # DOCKER_BUILDKIT=0 
	docker run --rm bcg-api make test-coverage

docs: # REQUIRES BUILDKIT
	docker build --target docs-output --output type=local,dest=src -t bcg-docs .


app: build # DOCKER_BUILDKIT=0 
	docker build --target app -t bcg-app . 

run: app # DOCKER_BUILDKIT=0 
docker run --rm -p 3000:3000 bcg-app
