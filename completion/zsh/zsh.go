package shells

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-go/utils/config"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io/ioutil"
	"path/filepath"
)

const zshAutocomplete = `_jfrog() {
	local -a opts
	opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
	_describe 'values' opts
	if [[ $compstate[nmatches] -eq 0 && $words[$CURRENT] != -* ]]; then
		_files
	fi
}

compdef _jfrog jfrog`

func WriteZshCompletionScript() {
	homeDir, err := config.GetJfrogHomeDir()
	if err != nil {
		log.Error(err)
		return
	}
	completionPath := filepath.Join(homeDir, "jfrog_zsh_completion")
	if err = ioutil.WriteFile(completionPath, []byte(zshAutocomplete), 0600); err != nil {
		log.Error(err)
		return
	}
	sourceCommand := "source " + completionPath + ""
	fmt.Printf(`Generated zsh completion script at %s.
To activate auto-completion on this shell only, source the completion script by running the following three commands:

autoload -Uz compinit
compinit
%s

To activate auto-completion permanently, put the above three commands in ~/.zshrc.

`,
		completionPath, sourceCommand)
}
