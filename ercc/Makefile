# Copyright 2019 Intel Corporation
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

TOP = ..
include $(TOP)/build.mk

all: build test docker

build: ercc

ercc: ercc_dependencies
# ERCC's binary is created here.
# The binary is then referenced by the fpc-peer cli, inserted in the ERCC package,
# and eventually run by the external launcher.
	$(GO) build $(GOTAGS) -o ercc main.go

ercc_dependencies:
# hard to list explicitly, so just leave empty target,
# which forces ecc to always be built

test:
	$(GO) test $(GOTAGS) -v ./...

clean: docker-clean
	$(GO) clean
	rm -f ercc

generate:
	$(GO) generate ./...

DOCKER_IMAGE?=fpc/ercc

# Note:
# - docker images are not necessarily rebuild if they exist but are outdated.
#   To force rebuild you have two options
#   - do a 'make clobber' first. This ensures you will have the uptodate images
#     but is a broad and slow brush
#   - to just fore rebuilding an image, call `make` with DOCKER_FORCE_REBUILD defined
#   - to keep docker build quiet unless there is an error, call `make` with DOCKER_QUIET_BUILD defined
DOCKER_BUILD_OPTS ?=
ifdef DOCKER_QUIET_BUILD
	DOCKER_BUILD_OPTS += --quiet
endif
ifdef DOCKER_FORCE_REBUILD
	DOCKER_BUILD_OPTS += --no-cache
endif
DOCKER_BUILD_OPTS += --build-arg FPC_VERSION=$(FPC_VERSION)

docker:
	$(DOCKER) build $(DOCKER_BUILD_OPTS) -t $(DOCKER_IMAGE) .
	$(DOCKER) tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):${FPC_VERSION}

docker-clean:
	@for tag in \
	    ${FPC_VERSION} \
	    latest \
	; do \
	    if [ ! -z "$$(docker images -q ${DOCKER_IMAGE}:$${tag})" ]; then \
			$(DOCKER) rmi $(DOCKER_IMAGE):$${tag}; \
	    fi \
	done

