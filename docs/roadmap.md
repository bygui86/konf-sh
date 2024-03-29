
# konf-sh - Roadmap

## Legend

| Emoji | Description               |
|:------|:--------------------------|
| ✅     | Released                  |
| 🏗    | In development            |
| 🗓    | Scheduled for development | 
| 🧐    | Evaluating                |

---

## v1.0.0 🗓

- [ ] create go install
- [ ] `TBD` create homebrew-tap (see https://goreleaser.com/customization/#Homebrew)
- [ ] have a look at [this library](https://github.com/gkarthiks/k8s-discovery)

## v0.8.0 🗓

- [ ] additional documentation
    - [ ] `TBD` go-doc
    - [ ] `TBD` add go-doc badge (see https://pkg.go.dev/github.com/etherlabsio/healthcheck?tab=doc)
- [ ] testing
    - [ ] unit tests
    - [ ] `TBD` add test badge (maybe https://github.com/bygui86/konf-sh/actions/workflows/test.yaml/badge.svg)
    - [ ] code-coverage check
    - [ ] `TBD` add code-coverage badge (see https://codecov.io/gh/etherlabsio/healthcheck)
- [ ] `TBD` introduce https://github.com/spf13/afero (A FileSystem Abstraction System for Go)

## v0.7.0 🗓

- [ ] implement namespaces commands
  - [ ] view current namespace
  - [ ] list all namespaces
  - [ ] set local (current shell only) namespace
  - [ ] set global namespace
  - [ ] set "-" to switch back namespace to previous
  - [ ] save last set namespace to "~/.kube/konfigs/last-ns"

## v0.6.0 🗓

- [ ] `TBD` put "global" as implicit
- [ ] implement "--silent" flag
- [ ] shellwrapper
  - [ ] implement command
  - [ ] fix "konf set local <context>"
  - [ ] fix "konf reset local"
- [ ] prompting
  - [ ] make it interactive with https://github.com/manifoldco/promptui

## v0.5.0 ✅

- [x] improve README and overall documentation
- [x] update codebase
  - [x] move to go 1.17
  - [x] libraries
  - [x] package structure
- [x] rename "configs" to "konfigs"
- [x] rename commands removing "-cfg" and "-ctx"
  - [x] clean/delete
  - [x] list
  - [x] rename
  - [x] reset
  - [x] set
  - [x] split
  - [x] view
- [x] rename binary to "konf-sh"
- [x] fix "list konfigs"
    - [x] hide .DS_store and other files
    - [x] show only valid kubeconfig
- [x] improve "clean/delete" command
  - [x] rename to "delete"
  - [x] remove ctx from both "~/.kube/config" and "~/.kube/konfigs" 
- [x] improve "rename" command
  - [x] rename ctx in both "~/.kube/config" and "~/.kube/konfigs" 
- [x] improve "split" command (see TODO in commands/set/action.go)
- [x] improve "completion" command
  - [x] fix "completion zsh" (replace "PROG" with "konf")
- [x] improve "set" command
  - [x] save last set ctx to "~/.kube/konfigs/last-ctx/"
  - [x] add "set -" to switch back ctx to previous
- [x] improve "list" command
  - [x] add listing ctx from "~/.kube/config"
- [x] review logging
  - [x] refactor logger embracing better and more standard approach with zap library
  - [x] review logs in all commands
    - [x] app
    - [x] completion
    - [x] clean/delete
    - [x] list
    - [x] rename
    - [x] reset
    - [x] set
    - [x] split
    - [x] view

## v0.4.0 ✅

- [x] fix GitHub Action release
- [x] fix Makefile bug in release target
- [x] improve release mechanism

## v0.3.0 ✅

- [x] add logo
- [x] align version everywhere

## v0.2.0 ✅

- [x] add arguments usage
- [x] new commands skeleton
- [x] skeleton for rename and reset commands
- [x] fix Makefile
- [x] implement second set of commands
- [x] improve clean command
- [x] improved config examples
- [x] improved logging
- [x] rename commands

## v0.1.0 ✅

- [x] implement first set of commands
- [x] implement properly logging flags
- [x] makefile
- [x] release mechanism
- [x] ci/cd with GitHub Actions
- [x] add 'ArgsUsage' description in all commands
- [x] improve release mechanism
