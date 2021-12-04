.PHONY: lint
lint:
	@golangci-lint run

.PHONY: build
build:
	@go build -o HiQidas ./*.go

.PHONY: run
run:
	@go run ./*.go

.PHONY: up
up:
	@cd dev && COMPOSE_PROJECT_NAME=hiqidas_hot_reload docker-compose up -d --build

.PHONY: down
down:
	@cd dev && docker-compose down

.PHONY: reset-frontend
reset-frontend: stop-front rm-front delete-front-image

.PHONY: stop-front
stop-front:
	@docker ps -a | grep HiQidas_frontend | awk '{print $$1}' | xargs docker stop

.PHONY: rm-front
rm-front:
	@docker ps -a | grep HiQidas_frontend | awk '{print $$1}' | xargs docker rm

.PHONY: delete-front-image
delete-front-image:
	@docker images -a | grep hiqidas | grep frontend | awk '{print $$3}' | xargs docker rmi

.PHONY: prune
prune:
	@docker image prune -a && docker volume prune

.PHONY: protobuf
protobuf:
	@mkdir -p server/protobuf
	@protoc --proto_path=protobuf --go_out=server/protobuf--go_opt=paths=source_relative protobuf/**/*.proto
	@protoc --doc_out=html,pb_rest.html:docs protobuf/rest/*.proto
	@protoc --doc_out=html,pb_ws.html:docs protobuf/ws/*.proto

.PHONY: tbls
tbls:
	@cd docs/db_schema && tbls doc

.PHONY: tbls-docker
tbls-docker:
	@docker exec -it -w /HiQidas $(docker ps -a | grep HiQidas_backend | awk '{print $$1}') make tbls

.PHONY: chown
chown:
	$(eval name := $(shell whoami))
	@sudo chown -R $(name):$(name) .
