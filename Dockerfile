#
#  Copyright Pnoker. All Rights Reserved.
#

FROM golang:1.13-alpine AS build

ENV GO111MODULE=on
WORKDIR /go/src/github.com/pnoker/go-project

COPY go.mod .

RUN go mod download

COPY . .

RUN cd cmd/ \
    && rm -rf http_server_mock.go \
    && go build -o go-project

FROM alpine AS base
MAINTAINER pnoker <pnokers.icloud.com>

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && mkdir -p /pnoker/go-project

WORKDIR /pnoker/go-project

COPY --from=build /go/src/github.com/pnoker/go-project/cmd/res/application-docker.toml ./res/application.toml
COPY --from=build /go/src/github.com/pnoker/go-project/cmd/go-project ./

CMD ./go-project
