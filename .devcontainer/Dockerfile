# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.3/containers/ubuntu/.devcontainer/base.Dockerfile

FROM namely/protoc-all:1.40_0

ENV TZ=Asia/Tokyo

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install curl sudo git vim mariadb-client -y --no-install-recommends

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/k1LoW/tbls@latest