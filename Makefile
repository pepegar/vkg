SHELL = /bin/bash

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
app = $(current_dir)


$(app): 
	@go build

test:
	@echo NO TESTS!

.PHONY: clean $(app) run release test
