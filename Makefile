SHELL = /usr/bin/env bash -o nounset -o pipefail -o errexit -c

ifneq ("$(wildcard .env-selected)","")
  env_file := $(shell cat .env-selected)

  ifneq ("$(wildcard $(env_file))","")
    include $(env_file)
    export
  endif
endif

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

lala: env_selected

.PHONY: env_selected
env_selected:
	@if [[ ! -f .env-selected ]]; then \
		echo "No .env-selected file found, please run 'make env'"; \
		exit 1; \
	fi; \
	\
	if [[ ! -f "$(shell cat .env-selected)" ]]; then \
		echo "$(shell cat .env-selected) does not exist, please either run 'make env dot=...' to select"; \
		echo "another file or create the file '$(shell cat .env-selected)' manually"; \
		exit 1; \
	fi

.PHONY: env
env:
	@if [[ -z "$(dot)" ]]; then \
		echo 'Specify the env-file you want to use'; \
		echo 'e.g.:'; \
		echo 'make env dot=dev'; \
		echo 'and make sure that the file .env.dev exists'; \
		exit 1; \
	fi

	@echo ".env.$(dot)" > .env-selected

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP_NAME) ./cmd/plannergen

