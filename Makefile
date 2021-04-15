# Copyright 2019 Intel Corporation
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

TOP = .
include $(TOP)/build.mk

SUB_DIRS = protos common internal ercc ecc_enclave ecc examples client_sdk utils fabric # docs
INTEGRATION_DIRS = integration

.PHONY: license


build: godeps

build test clean clobber:
	$(foreach DIR, $(SUB_DIRS), $(MAKE) -C $(DIR) $@ || exit;)

checks: linter license

license:
	@echo "License: Running licence checks.."
	@scripts/check_license.sh

linter: 
	@echo "LINT: Running code checks for Go files..."
	@cd $$(/bin/pwd) && ./scripts/golinter.sh
	@echo "LINT: Running code checks for Cpp/header files..."
	@cd $$(/bin/pwd) && ./scripts/cpplinter.sh
	@echo "LINT completed."

godeps: gotools
	$(GO) mod download

gotools:
	# install go tools if not present
	# (for faster docker-build, also replicte these commands
	#  in 'utils/docker/base-dev/Dockerfile')
	$(GO) install golang.org/x/tools/cmd/goimports
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go
	GO111MODULE=off $(GO) get github.com/maxbrunsfeld/counterfeiter
	GO111MODULE=on $(GO) get github.com/mikefarah/yq/v3

report:
	@echo "Reporting CI data..."
	@cd $$(/bin/pwd) && ./scripts/report.sh

integration-test:
	@echo "Running integration tests"
	$(MAKE) -C $(INTEGRATION_DIRS)

# add the ci_report target only at CI time, when coverage is enabled
ifeq (${IS_CI_RUNNING}, true)
ifeq (${CODE_COVERAGE_ENABLED}, true)
ci_report: report
endif
endif
