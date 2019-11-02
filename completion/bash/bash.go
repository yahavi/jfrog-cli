package shells

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-go/utils/config"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io/ioutil"
	"path/filepath"
)

const bashAutocomplete = `#!/bin/bash
_jfrog() {
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
}

complete -F _jfrog -o default jfrog
`

func WriteBashCompletionScript() {
	homeDir, err := config.GetJfrogHomeDir()
	if err != nil {
		log.Error(err)
		return
	}
	completionPath := filepath.Join(homeDir, "jfrog_bash_completion")
	if err = ioutil.WriteFile(completionPath, []byte(bashAutocomplete), 0600); err != nil {
		log.Error(err)
		return
	}
	sourceCommand := "source " + completionPath + ""
	fmt.Printf(`Generated bash completion script at %s.
To activate auto-completion on this shell only, source the completion script by running the following command:

%s

To activate auto-completion permanently, put the above command in ~/.bashrc or ~/.bash_profile, depending on your operating system.

`,
		completionPath, sourceCommand)
}
