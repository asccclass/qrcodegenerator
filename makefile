BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
APP?=app
PORT?=11005
ImageName?=sherry/qrcodegenerator
ContainerName?=qrcodegenerator
MKFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURDIR := $(dir $(MKFILE))

build:
	GOOS=linux GOARCH=amd64 go build -tags netgo \
	-ldflags "-s -w -X version.BuildTime=${BUILD_TIME}" \
	-o ${APP}

docker: build
	docker build -t ${ImageName} .
	rm -f app
	docker images


run: docker
	docker run --rm -d --name ${ContainerName} -v /etc/localtime:/etc/localtime:ro \
	-v ${CURDIR}/tmp:/app/tmp \
	--env-file ./envfile -p ${PORT}:80 ${ImageName}
	make log
	
stop:
	docker stop ${ContainerName}
	
log:
	docker logs -f -t --tail 20 ${ContainerName}

re:stop run
