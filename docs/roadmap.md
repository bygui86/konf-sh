
# konf - Roadmap

### v1.0.0 🗓

- [ ] create go install
- [ ] `TBD` create homebrew-tap (see https://goreleaser.com/customization/#Homebrew)
- [ ] have a look at [this library](https://github.com/gkarthiks/k8s-discovery)

### v0.7.0 🗓

- [ ] additional documentation
    - [ ] `TBD` go-doc
    - [ ] `TBD` add go-doc badge (see https://pkg.go.dev/github.com/etherlabsio/healthcheck?tab=doc)
- [ ] testing
    - [ ] unit tests
    - [ ] `TBD` add test badge (maybe https://github.com/bygui86/konf-sh/actions/workflows/test.yaml/badge.svg)
    - [ ] code-coverage check
    - [ ] `TBD` add code-coverage badge (see https://codecov.io/gh/etherlabsio/healthcheck)

### v0.6.0 🗓

- [ ] implement namespaces commands
  - [ ] view current namespace
  - [ ] list all namespaces
  - [ ] set local (current shell only) namespace
  - [ ] set global namespace

### v0.5.0 🏗

- [x] improve README and overall documentation
- [x] update codebase
  - [x] move to go 1.17
  - [x] libraries
  - [x] package structure
- [x] rename "configs" to "konfigs"
- [ ] rename commands removing "-cfg"
- [ ] improve "clean" command
  - [ ] rename to "delete"
  - [ ] remove ctx from both "~/.kube/config" and "~/.kube/konfigs" 
- [ ] improve "rename" command
  - [ ] rename ctx in both "~/.kube/config" and "~/.kube/konfigs" 
- [ ] improve split command (see TODO in commands/set/action.go)
- [ ] fix "completion zsh" (replace "PROG" with "konf")
- [ ] fix "list konfigs"
    - [ ] hide .DS_store and other files
    - [ ] show only valid kubeconfig
- [ ] rename binary to "konf-sh"
- [ ] improve "set"
  - [ ] add "-" to set "last konf used" in new shell 
- [ ] implement "--silent" flag
- [ ] shellwrapper
  - [ ] implement command
  - [ ] fix "konf set local <context>"

### v0.4.0 ✅

- [x] fix GitHub Action release
- [x] fix Makefile bug in release target
- [x] improve release mechanism

### v0.3.0 ✅

- [x] add logo
- [x] align version everywhere

### v0.2.0 ✅

- [x] add arguments usage
- [x] new commands skeleton
- [x] skeleton for rename and reset commands
- [x] fix Makefile
- [x] implement second set of commands
- [x] improve clean command
- [x] improved config examples
- [x] improved logging
- [x] rename commands

### v0.1.0 ✅

- [x] implement first set of commands
- [x] implement properly logging flags
- [x] makefile
- [x] release mechanism
- [x] ci/cd with GitHub Actions
- [x] add 'ArgsUsage' description in all commands
- [x] improve release mechanism
