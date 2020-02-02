
# konf
Kubernetes Configurator

`konf` makes easier to manage, maintain and use the Kubernetes configuration file (per default `~/.kube/config`).

---

## Status

![](https://github.com/bygui86/konf/workflows/build/badge.svg?branch=master)
&nbsp;&nbsp;&nbsp;![](https://github.com/bygui86/konf/workflows/release/badge.svg)

---

## Build

```shell
git clone git@github.com:bygui86/konf.git
cd konf
make build
```

---

## Commands

`konf split-cfg` separates the Kubernetes configuration (e.g. `~/.kube/config` if not otherwise specified) into single Kubernetes configurations files (per default saved in `~/.kube/configs/*`)

`konf list-cfg` lists all single Kubernetes configurations files separated by `konf split` (per default in `~/.kube/configs/*`)

`eval $(konf set-cfg local <context>)` sets the local (current shell) Kubernetes context (`KUBECONFIG` environment variable) to the specified one (per default `~/.kube/configs/*`) (*)

`konf set-cfg global <context>` sets the global Kubernetes context (per default in `~/.kube/config` Kubernetes configuration if not otherwise specified) to the specified one (per default `~/.kube/configs/*`)

`konf view-cfg` shows the local (current shell) and global Kubernetes context

`konf view-cfg local` shows only the local (current shell) Kubernetes context

`konf view-cfg global` shows only the global Kubernetes context

`konf clean-ctx <context-list_comma-separated>` removes the specified comma-separated context list from Kubernetes configuration (per default in `~/.kube/config` Kubernetes configuration if not otherwise specified)

`konf rename-ctx <context-to-rename> <new-context-name>` renames the specified context in Kubernetes configuration (per default in `~/.kube/config` Kubernetes configuration if not otherwise specified)

`eval $(konf reset-cfg local)` resets the local (current shell) Kubernetes configuration, un-setting `KUBECONFIG` environment variable

`konf reset-cfg global` resets `currentContext` to N/A in Kubernetes configuration (per default in `~/.kube/config` Kubernetes configuration if not otherwise specified)

`konf completion [bash | zsh]` outputs the auto-completion script for the selected. See [auto-completion](#auto-completion) section below.

`konf help` shows the helper

`konf version` shows the current version of konf

(*) INFO: The `konf set local` command must be executed in an `eval`, because it has to set the `KUBECONFIG` environment variable on the caller shell instance. 

---

## Makefile actions

### Build binary
```shell
make build
```

### Run
```shell
# from source
make run

# from bin
make run-bin
```

### Clean binary
```shell
make clean-bin
```

### Split a sample Kubernetes configuration file
```shell
make split-cfg
```

### List a set of sample Kubernetes configurations files
```shell
make list-cfg
```

### Set local Kubernetes context (current shell)
```shell
make set-cfg-local
```

### Set global Kubernetes context
```shell
make set-cfg-global
```

### View local and global Kubernetes contexts
```shell
make view-cfg
```

### View local Kubernetes context (current shell)
```shell
make view-cfg-local
```

### View global Kubernetes context
```shell
make view-cfg-global
```

### Clean Kubernetes contexts
```shell
make clean-ctx
```

### Rename Kubernetes context
```shell
make rename-ctx
```

### Release

`WARN`: Be careful, this command triggers the `release` GitHub Action that results in a new release on GitHub repo

```shell
make release NEW_VERSION=...
```

---

## Configurations

### Flags

| Flag | Command list | Available values | Default | Corresponding env-var | Description |
| --- | --- | --- | --- | --- | --- |
| --kube-config | split, view, view global, set global | - | $HOME/.kube/config | KONF_KUBE_CONFIG_PATH | Specify a custom Kubernetes configuration file path |
| --single-configs | split, list, set local, set global | - | $HOME/.kube/configs/ | KONF_SINGLE_KUBE_CONFIGS_PATH | Specify the single Kubernetes configurations files path |

### Environment variables

| Key | Command list | Available values | Default | Corresponding flag | Description |
| --- | --- | --- | --- | --- | --- |
| KONF_LOG_ENCODING | (global) | console, json | console | - | Set logger encoding |
| KONF_LOG_LEVEL | (global) | debug, info, warn, error, fatal | info | - | Set logger level |
| KONF_KUBE_CONFIG_PATH | split, view, view global, set global | - | $HOME/.kube/config | --kube-config | Specify a custom Kubernetes configuration file path |
| KONF_SINGLE_KUBE_CONFIGS_PATH | split, list, set local, set global | - | $HOME/.kube/configs/ | --single-configs | Specify the single Kubernetes configurations files path |

---

## Error codes

| Code | Command | Description |
| --- | --- | --- |
| 1 | (global) | Error initializing zap logger |
| 2 | (global) | Error starting application |
| 3 | (global) | Error creating specific application command |
| 11 | split-cfg | Error checking existence of Kubernetes configurations files path |
| 12 | split-cfg, set-cfg global, clean-ctx, rename-ctx | Error validating Kubernetes configuration (single, global, cleaned) |
| 13 | split-cfg, set-cfg global, clean-ctx, rename-ctx | Error writing Kubernetes configuration (single, global, cleaned) to file |
| 21 | list-cfg | Error listing single Kubernetes configurations |
| 31 | set-cfg local | Error checking existence of Kubernetes configurations files path |
| 32 | set-cfg local, set-cfg global | Error getting Kubernetes context: context argument not specified |
| 33 | set-cfg local | Error checking existence of Kubernetes context |
| 34 | set-cfg global, rename-ctx | Error checking existence of context in Kubernetes configuration |
| 41 | clean-ctx | Error getting Kubernetes context list: 'context list' argument not specified |
| 42 | clean-ctx | Error validating Kubernetes context list: 'context list' argument not valid. Context list must be a comma-separated list |
| 43 | clean-ctx | Error cleaning Kubernetes context list |
| 51 | rename-ctx | Error getting Kubernetes context to rename: 'context to rename' and 'new context name' arguments not specified |
| 52 | rename-ctx | Error getting Kubernetes context to rename: 'context to rename' argument not specified |
| 53 | rename-ctx | Error getting Kubernetes context to rename: 'new context name' argument not specified |
| 54 | rename-ctx | Error removing context from Kubernetes configuration |

---

## Auto-completion

### BASH

#### Method 1

Source the script file `commands/completion/bash_autocomplete` in your `.bashrc` or `.bash_profile` file.

#### Method 2

Execute following commands

```shell
echo 'source <(konf completion bash)' >> $HOME/.bashrc
. ./bashrc
konf
# now play with tab
```

### ZSH

#### Method 1

Take the script file `commands/completion/zsh_autocomplete`, replace `PROG` with `konf, source it `in your `.zshrc` file.

#### Method 2

Execute following commands

```shell
PROG=konf echo 'source <(konf completion zsh)' >> $HOME/.zshrc
konf
# now play with tab
```

---

## TODO list

- [x] implement initial commands
- [ ] implement additional commands
- [x] implement properly logging flags
- [x] documentation
- [x] makefile
- [ ] testing
- [x] release mechanism
- [x] ci/cd
- [ ] finalize split command (see TODO in commands/set/action.go)
- [x] add 'ArgsUsage' description in all commands

### Initial commands

- [x] split-cfg
- [x] list-cfg
- [x] set-cfg
- [x] view-cfg
- [x] help
- [x] version

### Additional commands

- [x] shell auto-completion
- [x] clean-ctx (remove specified context and relatives from kubernetes configuration)
- [x] rename-ctx (rename specified context in kubernetes configuration)
- [ ] reset-cfg
    - [ ] reset-cfg local removes KUBECONFIG environment variable
    - [ ] reset-cfg global resets currentContext to N/A in Kubernetes configuration

---

## Release

1. Choose a new version
	```shell
	NEW_VERSION="v0.1"
	```
2. Create a new tag with choosen version
	```shell
	git tag -a $NEW_VERSION -m "Tag for release $NEW_VERSION"
	```
3. Push new tag to remote, triggering `release` GitHub Action
	```shell
	git push origin $NEW_VERSION
	```

### Available mechanisms

- goreleaser
- GitHub Actions
- GitHub Package Registry
- PackagePublishing

### GitHub Actions

| Action | Triggered by | Steps |
| --- | --- | --- |
| build | push to master, push to branch features/\*\*, PR to master, PR to branch features/\*\* | setup go, checkout, get dependencies, build, test |
| release | new tag creation | setup go, checkout, unshallow, run goreleaser |

### goreleaser

`WARN`: The first three steps will trigger the `release` GitHub Action, performing the last step (goreleaser), so be careful if you want to release manually.

1. version
	```shell
	NEW_VERSION="v0.1"
	```
2. tag
	```shell
	git tag -a $NEW_VERSION -m "Tag for release $NEW_VERSION"
	```
3. push
	```shell
	git push origin $NEW_VERSION
	```
4. release
	```shell
	goreleaser release --rm-dist
	```

---

## Links

### Golang
- https://github.com/golang/go/wiki/Modules
#### Logger
- https://github.com/uber-go/zap
- https://github.com/sandipb/zap-examples

### Kubernetes
- https://github.com/kubernetes/kubernetes/blob/master/staging/README.md
#### client-go
- https://godoc.org/k8s.io/client-go
- https://github.com/kubernetes/client-go/blob/master/INSTALL.md#add-client-go-as-a-dependency
- https://github.com/kubernetes/client-go/
- https://github.com/kubernetes/client-go/tree/master/examples
#### examples
- https://rancher.com/using-kubernetes-api-go-kubecon-2017-session-recap/
- https://github.com/alena1108/kubecon2017/blob/master/main.go

### Release

#### goreleaser
- https://github.com/bygui86/go-releaser
- https://goreleaser.com/
- https://goreleaser.com/actions/
- https://github.com/goreleaser/goreleaser-action
- https://github.com/marketplace/actions/goreleaser-action

#### GitHub Actions
- https://help.github.com/en/articles/about-github-actions
- https://help.github.com/en/articles/configuring-a-workflow
- https://help.github.com/en/articles/workflow-syntax-for-github-actions
- https://github.com/actions/setup-go

#### GitHub Package Registry
- https://github.com/features/packages

#### (Golang) PackagePublishing
- https://github.com/golang/go/wiki/PackagePublishing
