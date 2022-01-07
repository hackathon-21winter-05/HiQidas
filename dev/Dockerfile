FROM golang:1.17.6-alpine AS builder

RUN go install github.com/jwilder/dockerize@latest
RUN go install github.com/cosmtrek/air@latest

FROM golang:1.17.6-alpine AS hotReload
WORKDIR /HiQidas

COPY --from=builder /go/bin/dockerize /usr/local/bin/
COPY --from=builder /go/bin/air /usr/local/bin/

EXPOSE 7070

ENTRYPOINT dockerize -timeout 60s -wait tcp://mariadb:3306 air
