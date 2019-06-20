-include .env
VERSION=dev
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build .

GOCMD=go
BINARY=./dist/pg_reloaded

all: build
build: deps
	$(GOCMD) build -o $(BINARY)
deps:
	$(GOCMD) list -m all

docker:
	for i in 9.2 9.3 9.4 9.5 9.6 10.0 10.1 10.2 10.3 10.4 10.5 10.6 10.7 10.8 11.0 11.1 11.2 11.3; \
	do \
		$(DOCKERBUILD) --build-arg IMAGE=postgres:$$i-alpine -t gkawamoto/pg_reloaded:$$i-$(VERSION)-alpine; \
	done
