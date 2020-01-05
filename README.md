
# konf
Kubernetes Configurator

`konf` makes easier to manage, maintain and use the Kubernetes configuration file (per default `~/.kube/config`).

---

## Run

### from source
```shell
go run main.go
```

### from bin
```shell
konf ...
```

---

## Commands

- [ ] split
- [ ] list
- [ ] set
    - [ ] global
    - [ ] local
- [ ] help
- [ ] completion

---

## Configurations

### Flags

`TBD`

| Flag | Available values | Default |
| --- | --- | --- |
| kubeconfig | - | $HOME/.kube/config |

### Environment variables

`TBD`

| Key | Available values | Default |
| --- | --- | --- |
| LOG_ENCODING | console, json | console |
| LOG_LEVEL | debug, info, warn, error, dpanic, panic, fatal | info |
| KUBECONFIG | - | (empty) |

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
