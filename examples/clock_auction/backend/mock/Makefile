# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

TOP = ../../../..
include $(TOP)/build.mk

DOCKER_IMAGE = auction_mock_backend
APP_SRC = ${PWD}

build:
	$(DOCKER) build -t $(DOCKER_IMAGE) .

run:
	$(DOCKER) run \
		-it \
		-v ${APP_SRC}:/usr/src/app \
		-p 3000:3000 \
		$(DOCKER_IMAGE)
