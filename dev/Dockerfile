FROM golang:1.17.4-alpine
WORKDIR /HiQidas

RUN go install github.com/jwilder/dockerize@latest
RUN go install github.com/cosmtrek/air@latest

EXPOSE 7070

ENTRYPOINT dockerize -timeout 60s -wait tcp://mariadb:3306 air
