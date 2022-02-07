
# VARIABLES
KONF_LOG_LEVEL ?= info
KONF_PREFIX := KONF_LOG_LEVEL=$(KONF_LOG_LEVEL)


# CONFIG
.PHONY: build clean debug help list print-variables release run set-global set-local split view view-global view-local
.DEFAULT_GOAL := help


# ACTIONS

## commands

run :		## Debug running directly from source code
	go run main.go

build :		## Build
	@export GO111MODULE=on && \
	go build -o konf-sh .

run-bin : build		## Run
	konf-sh $(ARGS)

clean-bin : 		## Clean binary
	@rm -rf konf-sh >/dev/null 2>&1

release :		## Create a new git tag and push it to remote to trigger the release GitHub action
ifdef NEW_VERSION
	git tag -a $(NEW_VERSION) -m "Tag for release $(NEW_VERSION)"
	git push origin $(NEW_VERSION)
else
	@echo "New version (environment variable NEW_VERSION) not specified, please run command again like 'make release NEW_VERSION=...'"
endif

simulate-release :		## Simulate a release with goreleaser
	goreleaser release --rm-dist --snapshot --skip-publish

## features samples

split : build		## Split a sample Kubernetes configuration file
	$(KONF_PREFIX) konf-sh split --kube-config ./examples/config --single-konfigs ./examples/konfigs

list : build		## List a set of sample Kubernetes konfigurations files
	$(KONF_PREFIX) konf-sh list --single-konfigs ./examples/konfigs

set-local : build		## Set local Kubernetes context (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf-sh set local context_b --single-konfigs ./examples/konfigs)'"
	@echo ""

set-global : build		## Set global Kubernetes context
	$(KONF_PREFIX) konf-sh set global context_b --kube-config ./examples/config

view : build		## View local and global Kubernetes contexts
	$(KONF_PREFIX) konf-sh view --kube-config ./examples/config

view-local : build		## View local Kubernetes context (current shell)
	$(KONF_PREFIX) konf-sh view local

view-global : build		## View global Kubernetes context
	$(KONF_PREFIX) konf-sh view global --kube-config ./examples/config

delete : build		## Remove context list from Kubernetes configuration
	$(KONF_PREFIX) konf-sh delete --kube-config ./examples/config context_a,context_b

rename : build		## Rename specified context in Kubernetes configuration
	$(KONF_PREFIX) konf-sh rename --kube-config ./examples/config context_a NEW_context_a

reset-local : build		## Reset local Kubernetes configuration (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf-sh reset local)'"
	@echo ""

reset-global : build		## Reset global Kubernetes configuration
	$(KONF_PREFIX) konf-sh reset global --kube-config ./examples/config

## helpers

reset-config-sample :		## Reset Kubernetes configuration sample to original
	cp -f ./examples/config_origin ./examples/config

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo ""
