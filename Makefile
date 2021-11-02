-include .env
MODIFIER=-master
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build .
DOCKERPUSH=$(DOCKERCMD) push
DOCKERTAG=$(DOCKERCMD) tag

GOCMD=go
BINARY=./dist/pg_reloaded

all: build
build: deps
	$(GOCMD) build -o $(BINARY)
deps:
	$(GOCMD) list -m all

docker:
	COMMIT=$$(git rev-parse --short HEAD); \
	for i in 9.2 9.3 9.4 9.5 9.6 10 11 12 13 14; \
	do \
		$(DOCKERBUILD) --build-arg IMAGE=postgres:$$i-alpine -t gkawamoto/pg_reloaded:$$i$(MODIFIER)-alpine; \
		$(DOCKERTAG) gkawamoto/pg_reloaded:$$i$(MODIFIER)-alpine gkawamoto/pg_reloaded:$$COMMIT-$$i$(MODIFIER)-alpine; \
	done
docker-publish: docker
	COMMIT=$$(git rev-parse --short HEAD); \
	for i in 9.2 9.3 9.4 9.5 9.6 10 11 12 13 14; \
	do \
		$(DOCKERPUSH) gkawamoto/pg_reloaded:$$i$(MODIFIER)-alpine; \
		$(DOCKERPUSH) gkawamoto/pg_reloaded:$$COMMIT-$$i$(MODIFIER)-alpine; \
	done
