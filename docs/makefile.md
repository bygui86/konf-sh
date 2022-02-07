
# konf - Makefile targets

### Build binary

```sh
make build
```

### Run

```sh
# from source
make run

# from bin
make run-bin
```

### Clean binary

```sh
make clean-bin
```

### Split a sample Kubernetes configuration file

```sh
make split-cfg
```

### List a set of sample Kubernetes configurations files

```sh
make list-cfg
```

### Set local Kubernetes context (current bash)

```sh
make set-cfg-local
```

### Set global Kubernetes context

```sh
make set-cfg-global
```

### View local and global Kubernetes contexts

```sh
make view-cfg
```

### View local Kubernetes context (current bash)

```sh
make view-cfg-local
```

### View global Kubernetes context

```sh
make view-cfg-global
```

### Clean Kubernetes contexts

```sh
make clean-ctx
```

### Rename Kubernetes context

```sh
make rename-ctx
```

### Release

`WARN`: Be careful, this command triggers the `release` GitHub Action that results in a new release on GitHub repo

```sh
make release NEW_VERSION=...
```
