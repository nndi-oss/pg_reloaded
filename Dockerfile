ARG IMAGE=postgres:11.2-alpine

FROM golang:1.14-alpine AS builder
RUN apk add --no-cache git make
ENV GO111MODULE on
COPY . /go/src/github.com/nndi-oss/pg_reloaded/
RUN cd /go/src/github.com/nndi-oss/pg_reloaded/ && make build BINARY=/dist/pg_reloaded

FROM $IMAGE
COPY --from=builder /dist/pg_reloaded /usr/bin/pg_reloaded
CMD ["/usr/bin/pg_reloaded", "start", "--config", "/etc/pg_reloaded/pg_reloaded.yml"]
