package completion

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const (
	bashScript = `#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete $PROG
unset PROG`

	zshScript = `#compdef konf

_cli_zsh_autocomplete() {

  local -a opts
  opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")

  _describe 'values' opts

  return
}

compdef _cli_zsh_autocomplete konf`
)

func bashCompletion(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing BASH-COMPLETION command")

	zap.L().Info(bashScript)
	return nil
}

func zshCompletion(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing ZSH-COMPLETION command")

	zap.L().Info(zshScript)
	return nil
}
