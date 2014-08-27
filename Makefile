SHELL = /bin/bash
PACKAGES := ./commands ./config/vkgrc

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
app = $(current_dir)


$(app):
	@go build

test:
	@go test $(PACKAGES)

.PHONY: clean $(app) run release test
