# sectfs - Makefile
# Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
#
# Use of this source code is governed by a BSD 2-Clause license that can be
# found in the LICENSE file.

APP := sectfs
DIR := bin

.PHONY: all clean build linux

all: linux

clean:
	@if [ -d ${DIR} ]; then rm -rf ${DIR}/*; else exit 0; fi

build:
	go build -o ${DIR}/${APP} -ldflags "-s -w" .

linux:
	@# linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${DIR}/${APP} -ldflags "-s -w" .
	@cd ${DIR}/ && mkdir config static ${APP}-server && cp ../config/sectfs.conf config/ && cp ../static/index.html static/
	@cd ${DIR}/ && mv ${APP} ${APP}-server/ && mv config/ ${APP}-server/ && mv static/ ${APP}-server/
	@cd ${DIR}/ && tar -zcf ${APP}-server-linux_amd64.tar.gz ${APP}-server/ && rm -rf ${APP}-server/ && cd ../
	@# linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${DIR}/${APP} -ldflags "-s -w" .
	@cd ${DIR}/ && mkdir config static ${APP}-server && cp ../config/sectfs.conf config/ && cp ../static/index.html static/
	@cd ${DIR}/ && mv ${APP} ${APP}-server/ && mv config/ ${APP}-server/ && mv static/ ${APP}-server/
	@cd ${DIR}/ && tar -zcf ${APP}-server-linux_arm64.tar.gz ${APP}-server/ && rm -rf ${APP}-server/ && cd ../
