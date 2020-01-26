
# konf
Kubernetes Configurator

`konf` makes easier to manage, maintain and use the Kubernetes configuration file (per default `~/.kube/config`).

---

## Build
```shell
git clone git@github.com:bygui86/konf.git
cd konf
make build
```

---

## Run
```shell
# from source
make debug

# from bin
make run
```

### Split a sample Kubernetes configuration file
```shell
make split
```

### List a set of sample Kubernetes configurations files
```shell
make list
```

### Set local Kubernetes context (current shell)
```shell
make set-local
```

### Set global Kubernetes context
```shell
make set-global
```

### View local and global Kubernetes contexts
```shell
make view
```

### View local Kubernetes context (current shell)
```shell
make view-local
```

### View global Kubernetes context
```shell
make view-global
```

---

## Commands

`konf split` separates the Kubernetes configuration (e.g. `~/.kube/config` if not otherwise specified) into single Kubernetes configurations files (per default saved in `~/.kube/configs/*`)

`konf list` lists all single Kubernetes configurations files separated by `konf split` (per default in `~/.kube/configs/*`)

`eval $(konf set local <context>)` sets the local (current shell) Kubernetes context to the specified one (per default `~/.kube/configs/*`) (*)

`konf set global <context>` sets the global Kubernetes context (per default in `~/.kube/config` Kubernetes configuration if not otherwise specified) to the specified one (per default `~/.kube/configs/*`)

`konf view` shows the local (current shell) and global Kubernetes context

`konf view local` shows only the local (current shell) Kubernetes context

`konf view global` shows only the global Kubernetes context

`konf completion <bash|zsh>` outputs the auto-completion script for the selected. See [auto-completion](#auto-completion) section below.

`konf help` shows the helper

`konf version` shows the current version of konf

(*) INFO: The `konf set local` command must be executed in an `eval`, because it has to set the `KUBECONFIG` environment variable on the caller shell instance. 

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
| 11 | split | Error checking existence of Kubernetes configurations files path |
| 12 | split, set global | Error validating single Kubernetes configuration |
| 13 | split, set global | Error writing single Kubernetes configuration to file |
| 21 | list | Error listing single Kubernetes configurations |
| 31 | set local | Error checking existence of Kubernetes configurations files path |
| 32 | set local, set global | Error getting Kubernetes context: context argument not specified |
| 33 | set local | Error checking existence of Kubernetes context |
| 34 | set global | Error checking existence of context in Kubernetes configuration |

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
- [ ] release mechanism
- [ ] ci/cd

### Initial commands

- [x] split
- [x] list
- [x] set
- [x] view
- [x] help
- [x] version

### Additional commands

- [x] shell auto-completion
- [ ] reset (reset local to defualt ~/.kube/config)
- [ ] clean kube-config (remove specified context from ~/.kube/config)
- [ ] rename context

---

## Release mechanisms

### Options

goreleaser
- https://github.com/bygui86/go-releaser
- https://goreleaser.com/quick-start/
- https://goreleaser.com/templates/

GitHub Actions
- https://help.github.com/en/articles/about-github-actions
- https://help.github.com/en/articles/configuring-a-workflow
- https://help.github.com/en/articles/workflow-syntax-for-github-actions
- https://github.com/actions/setup-go

GitHub Package Registry
- https://github.com/features/packages

PackagePublishing
- https://github.com/golang/go/wiki/PackagePublishing

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
