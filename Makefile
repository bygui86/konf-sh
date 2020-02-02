
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
	go build -o konf .

run-bin : build		## Run
	konf $(ARGS)

clean-bin : 		## Clean binary
	@rm -rf konf >/dev/null 2>&1

release :		## Create a new git tag and push it to remote to trigger the release GitHub action
ifdef NEW_VERSION
	git tag -a $(NEW_VERSION) -m "Tag for release $(NEW_VERSIO)"
	git push origin $(NEW_VERSION)
else
	@echo "New version (environment variable NEW_VERSION) not specified, please run command again like 'make release NEW_VERSION=...'"
endif

## features samples

split-cfg : build		## Split a sample Kubernetes configuration file
	$(KONF_PREFIX) konf split --kube-config ./examples/config --single-configs ./examples/configs

list-cfg : build		## List a set of sample Kubernetes configurations files
	$(KONF_PREFIX) konf list --single-configs ./examples/configs

set-cfg-local : build		## Set local Kubernetes context (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf set local context_b --single-configs ./examples/configs)'"
	@echo ""

set-cfg-global : build		## Set global Kubernetes context
	$(KONF_PREFIX) konf set global context_b --kube-config ./examples/config

view-cfg : build		## View local and global Kubernetes contexts
	$(KONF_PREFIX) konf view --kube-config ./examples/config

view-cfg-local : build		## View local Kubernetes context (current shell)
	$(KONF_PREFIX) konf view local

view-cfg-global : build		## View global Kubernetes context
	$(KONF_PREFIX) konf view global --kube-config ./examples/config

clean-ctx : build		## Remove context list from Kubernetes configuration
	$(KONF_PREFIX) konf clean --kube-config ./examples/config context_a,context_b

rename-ctx : build		## Rename specified context in Kubernetes configuration
	$(KONF_PREFIX) konf rename --kube-config ./examples/config context_a NEW_context_a

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
