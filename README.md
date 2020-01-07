
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

- [x] split
- [x] list
- [ ] set
    - [ ] local
    - [ ] global
- [x] view
    - [x] all
    - [x] local
    - [x] global
- [x] completion
- [x] help
- [x] version

---

## Configurations

### Flags `TO BE UPDATED`

| Flag | Command list | Available values | Default | Corresponding env-var | Description |
| --- | --- | --- | --- | --- | --- |
| --kube-config | split | - | $HOME/.kube/config | KONF_KUBE_CONFIG_PATH | Specify a custom Kubernetes configuration file path |
| --single-configs | split, list | - | $HOME/.kube/configs/ | KONF_SINGLE_KUBE_CONFIGS_PATH | Specify the single Kubernetes configurations files path |

### Environment variables `TO BE UPDATED`

| Key | Command list | Available values | Default | Corresponding flag | Description |
| --- | --- | --- | --- | --- | --- |
| KONF_LOG_ENCODING | (global) | console, json | console | - | Set logger encoding |
| KONF_LOG_LEVEL | (global) | debug, info, warn, error, fatal | info | - | Set logger level |
| KONF_KUBE_CONFIG_PATH | split | - | $HOME/.kube/config | --kube-config | Specify a custom Kubernetes configuration file path |
| KONF_SINGLE_KUBE_CONFIGS_PATH | split, list | - | $HOME/.kube/configs/ | --single-configs | Specify the single Kubernetes configurations files path |

---

## Error codes

| Code | Command | Description |
| --- | --- | --- |
| 1 | (global) | Error initializing zap logger |
| 2 | (global) | Error starting application |
| 3 | (global) | Error creating specific application command |
| 11 | split | Error checking existence of Kubernetes configurations files path |
| 12 | split | Error validating single Kubernetes configuration |
| 13 | split | Error writing single Kubernetes configuration to file |
| 21 | list | Error listing single Kubernetes configurations |
| 31 | set local | Error checking existence of Kubernetes configurations files path |
| 32 | set local | Error getting Kubernetes context: context argument not specified |
| 33 | set local | Error checking existence of Kubernetes context |
| 34 | set local | Error setting local Kubernetes context (env-var KUBECONFIG) |

---

## Auto-completion

### BASH

Source the `commands/completion/bash_autocomplete` file in your `.bashrc` or `.bash_profile` file. 

```shell
go build -o konf .
source <(konf completion bash)
konf
# now play with tab
```

### ZSH

Source the `commands/completion/zsh_autocomplete` file in your `.zshrc` file, while setting the `PROG` variable to the name of your program.

```shell
go build -o konf .
PROG=konf source <(konf completion zsh)
konf
# now play with tab
```

---

## TODO list

- [ ] implement commands
- [ ] implement properly logging flags
- [ ] documentation
- [x] makefile
- [ ] testing
- [ ] `TBD` container (see `hadolint` as example)

---

## Links

## Golang
- https://github.com/golang/go/wiki/Modules
### Logger
- https://github.com/uber-go/zap
- https://github.com/sandipb/zap-examples

## Kubernetes
- https://github.com/kubernetes/kubernetes/blob/master/staging/README.md
### client-go
- https://godoc.org/k8s.io/client-go
- https://github.com/kubernetes/client-go/blob/master/INSTALL.md#add-client-go-as-a-dependency
- https://github.com/kubernetes/client-go/
- https://github.com/kubernetes/client-go/tree/master/examples
### examples
- https://rancher.com/using-kubernetes-api-go-kubecon-2017-session-recap/
- https://github.com/alena1108/kubecon2017/blob/master/main.go
