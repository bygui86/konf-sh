
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
make split
```

### List a set of sample Kubernetes konfigurations

```sh
make list
```

### Set local Kubernetes context (current bash)

```sh
make set-local
```

### Set global Kubernetes context

```sh
make set-global
```

### View local and global Kubernetes contexts

```sh
make view
```

### View local Kubernetes context (current bash)

```sh
make view-local
```

### View global Kubernetes context

```sh
make view-global
```

### Clean Kubernetes contexts

```sh
make delete
```

### Rename Kubernetes context

```sh
make rename
```

### Release

`WARN`: Be careful, this command triggers the `release` GitHub Action that results in a new release on GitHub repo

```sh
make release NEW_VERSION=...
```
