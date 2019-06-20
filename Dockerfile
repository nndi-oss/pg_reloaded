ARG IMAGE=postgres:11.2-alpine

FROM golang:1.12-alpine AS builder
RUN apk add --no-cache git
ENV GO111MODULE on
COPY . /go/src/github.com/zikani03/pg_reloaded/
RUN cd /go/src/github.com/zikani03/pg_reloaded/ && go build -o /dist/pg_reloaded

FROM $IMAGE
COPY --from=builder /dist/pg_reloaded /usr/bin/pg_reloaded
CMD ["/usr/bin/pg_reloaded", "start", "--config", "/etc/pg_reloaded/pg_reloaded.yml"]
