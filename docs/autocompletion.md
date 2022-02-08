
# konf-sh - Autocompletion

Current supported shells: zsh, bash

## BASH

### Method 1

1. source the script file `commands/completion/bash_autocomplete` in your `.bashrc` or `.bash_profile` file

### Method 2 `[recommended]`

Execute following commands

```sh
echo 'source <(konf completion bash)' >> $HOME/.bashrc

. ./bashrc

konf

# now play with tab
```

---

## ZSH

### Method 1

1. take the script file `commands/completion/zsh_autocomplete`
2. replace `PROG` with `konf`
3. source it in your `.zshrc` file

### Method 2 `[recommended]`

Execute following commands

```sh
# OLD, deprecated, to be replaced soon
PROG=konf echo 'source <(konf completion zsh)' >> $HOME/.zshrc
# NEW, recommended
echo 'source <(konf completion zsh)' >> $HOME/.zshrc

konf

# now play with tab
```
