
# VARIABLES
KONF_LOG_ENCODING ?= console
KONF_LOG_LEVEL ?= info
KONF_PREFIX := KONF_LOG_ENCODING=$(KONF_LOG_ENCODING) KONF_LOG_LEVEL=$(KONF_LOG_LEVEL)
## global
export GO111MODULE = on


# CONFIG
.PHONY: build clean debug help list print-variables release run set-global set-local split view view-global view-local
.DEFAULT_GOAL := help


# ACTIONS

## code

build :		## Build
	go build -o konf-sh .

clean :		## Clean
	-@rm -rf konf-sh

test :		## Test
	go test -coverprofile=coverage.out -count=5 -race ./...

## release

release :		## Create a new git tag and push it to remote to trigger the release GitHub action
ifdef NEW_VERSION
	git tag -a $(NEW_VERSION) -m "Tag for release $(NEW_VERSION)"
	git push origin $(NEW_VERSION)
else
	@echo "New version (environment variable NEW_VERSION) not specified, please run command again like 'make release NEW_VERSION=...'"
endif

simulate-release :		## Simulate a release with goreleaser
	goreleaser release --rm-dist --snapshot --skip-publish

## commands

split : build		## Split a sample Kubernetes configuration file
	$(KONF_PREFIX) konf-sh split --kube-config ./examples/config --single-konfigs ./examples/konfigs

list : build		## List a set of sample Kubernetes konfigurations
	$(KONF_PREFIX) konf-sh list --kube-config ./examples/config --single-konfigs ./examples/konfigs

# TODO refactor after shell wrapper implementation
set-local : build		## Set local Kubernetes context (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf-sh set local --single-konfigs ./examples/konfigs context_b)'"
	@echo ""

# TODO refactor after shell wrapper implementation
set-last-local : build		## Set last Kubernetes context as local (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf-sh set local --single-konfigs ./examples/konfigs context_b)'"
	@echo ""

set-global : build		## Set global Kubernetes context
	$(KONF_PREFIX) konf-sh set global --kube-config ./examples/config context_b

set-last-global : build		## Set last Kubernetes context as global
	$(KONF_PREFIX) konf-sh set global --kube-config ./examples/config --single-konfigs ./examples/konfigs -

view : build		## View local and global Kubernetes contexts
	$(KONF_PREFIX) konf-sh view --kube-config ./examples/config

view-local : build		## View local Kubernetes context (current shell)
	$(KONF_PREFIX) konf-sh view local

view-global : build		## View global Kubernetes context
	$(KONF_PREFIX) konf-sh view global --kube-config ./examples/config

delete : build		## Remove context list from Kubernetes configuration
	$(KONF_PREFIX) konf-sh delete --kube-config ./examples/config --single-konfigs ./examples/konfigs context_a,context_b

rename : build		## Rename specified context in Kubernetes configuration
	$(KONF_PREFIX) konf-sh rename --kube-config ./examples/config --single-konfigs ./examples/konfigs context_a context_x

# TODO refactor after shell wrapper implementation
reset-local : build		## Reset local Kubernetes configuration (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf-sh reset local)'"
	@echo ""

reset-global : build		## Reset global Kubernetes configuration
	$(KONF_PREFIX) konf-sh reset global --kube-config ./examples/config

## helpers

restore-origin-example :		## Restore original example
	@rm -rf ./examples/konfigs/context_*
	@cp -f ./examples/config_origin ./examples/config
	@cp -f ./examples/context_invalid ./examples/konfigs/context_invalid

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
