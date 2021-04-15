# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

TOP = ..
include $(TOP)/build.mk

FABS := $(TOP)/fabric/_internal

build: $(FABS)

$(FABS):
	mkdir -p $@ && cd $@ && \
	curl -sSL https://bit.ly/2ysbOFE | bash -s -- $(FABRIC_VERSION) 1.4.9 -s -d

clean:
	rm -rf $(FABS)
