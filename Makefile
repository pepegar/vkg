SHELL = /bin/bash

test:
	@go test ./...

goveralls:
	@go test -coverprofile commands.part ./commands
	@go test -coverprofile vkgrc.part ./config/vkgrc
	@echo "mode: set" >coverage.out
	@grep -h -v "mode: set" *.part >>coverage.out
	@rm *.part
	@gocov convert coverage.out > coverage.json
	@goveralls -repotoken $(COVERALLS_REPO_TOKEN) -gocovdata="coverage.json"

open_cov:
	@go test -coverprofile commands.part ./commands
	@go test -coverprofile vkgrc.part ./config/vkgrc
	@echo "mode: set" >coverage.out
	@grep -h -v "mode: set" *.part >>coverage.out
	@rm *.part
	@go tool cover -html=coverage.out
	@git clean -fd
