.ONESHELL:
.DELETE_ON_ERROR:

# h - help
h help:
	@echo "h help 	- this help"
	@echo "build 	- run build docker image for api"
.PHONY: h

build:
	docker build -f ./build/restapi/Dockerfile -t plantbook/restapi .
.PHONY: build

builddb:
	docker build -f ./build/db/Dockerfile -t plantbook/database .
.PHONY: builddb

run:
	docker run --rm -it --name plantbook_api -p 8081:8081 plantbook/restapi
.PHONY: run

rundb:
	docker run -d --name plantbook_db -p 54321:5432 plantbook/database
.PHONY: rundb

publish:
	echo $$CR_PAT | docker login ghcr.io --username veremchukvv --password-stdin
	docker image tag plantbook/restapi ghcr.io/veremchukvv/plantbook/restapi:latest
	docker image push ghcr.io/veremchukvv/plantbook/restapi:latest
.PHONY: publish