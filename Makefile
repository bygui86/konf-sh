
# VARIABLES
# -


# CONFIG
.PHONY: build debug help list print-variables run set-global set-local split view view-global view-local
.DEFAULT_GOAL := help


# ACTIONS
build :		## Build
	@export GO111MODULE=auto && \
	go build -o konf .

run : build		## Run
	konf

debug :		## Debug running directly from source code
	go run main.go

split : build		## Split a sample Kubernetes configuration file
	konf split --kube-config ./examples/config --single-configs ./examples/configs

list : build		## List a set of sample Kubernetes configurations files
	konf list --single-configs ./examples/configs

set-local : build		## Set local Kubernetes context (current shell)
	@echo "It's useless to run an 'eval' command from the Makefile as each line is executed in a new shell instance"
	@echo "Please manually execute 'eval $(konf set local context_b --single-configs ./examples/configs)'"
	@echo ""

set-global : build		## Set global Kubernetes context
	@echo "Work in progress!"
	@echo ""

view : build		## View local and global Kubernetes contexts
	konf view --kube-config ./examples/config

view-local : build		## View local Kubernetes context (current shell)
	konf view local

view-global : build		## View global Kubernetes context
	konf view global --kube-config ./examples/config

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
