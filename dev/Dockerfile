FROM namely/protoc-all:1.40_0
WORKDIR /HiQidas

RUN go install github.com/jwilder/dockerize@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/k1LoW/tbls@latest

EXPOSE 7070

ENTRYPOINT dockerize -timeout 60s -wait tcp://mariadb:3306 air