
# VARIABLES
# -


# CONFIG
.PHONY: build debug help list print-variables run set-global set-local split view-global view-local
.DEFAULT_GOAL := help


# ACTIONS
build :		## Build
	@export GO111MODULE=auto && \
	go build -o konf .

run : build		## Run
	@export GO111MODULE=auto && \
	konf

debug :		## Debug running directly from source code
	@export GO111MODULE=auto && \
	go run main.go

# split : build		## Split a sample Kubernetes configuration file
split :		## Split a sample Kubernetes configuration file
	@export GO111MODULE=auto && \
	konf split --kube-config ./examples/config --single-configs ./examples/configs

# list : build split		## List a set of sample Kubernetes configurations files
list :		## List a set of sample Kubernetes configurations files
	@export GO111MODULE=auto && \
	konf list --single-configs ./examples/configs

# set-local : build split		## Set local Kubernetes context (current shell)
set-local :		## Set local Kubernetes context (current shell)
	@export GO111MODULE=auto && \
	konf set local context_b --single-configs ./examples/configs

# set-global : build split		## Set global Kubernetes context
set-global :		## Set global Kubernetes context
	@echo "Work in progress!"
	@echo ""

# view : build split set-local set-global		## View local and global Kubernetes contexts
view :		## View local and global Kubernetes contexts
	@export GO111MODULE=auto && \
	konf view --kube-config ./examples/config

# view-local : build split set-local		## View local Kubernetes context (current shell)
view-local :		## View local Kubernetes context (current shell)
	@export GO111MODULE=auto && \
	konf view local

# view-global : build split set-global		## View global Kubernetes context
view-global :		## View global Kubernetes context
	@export GO111MODULE=auto && \
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
