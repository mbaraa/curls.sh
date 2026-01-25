package scripts

import (
	"fmt"
	"strings"
	"text/template"
)

// bashScripts holds current bash scripts, each script is appended in its own file in this package.
var bashScripts []bashScript

type bashScript struct {
	id          string
	description string
	scriptText  string
}

func (s bashScript) Id() string {
	return s.id
}

func (s bashScript) Description() string {
	return s.description
}

func (s bashScript) RunString(args []string) string {
	passedArgs := make(map[string]string)
	for i, arg := range args {
		passedArgs[fmt.Sprintf("V%d", i+1)] = arg
	}

	out := new(strings.Builder)
	t := template.Must(template.New("script").Parse(s.scriptText))
	err := t.Execute(out, passedArgs)
	if err != nil {
		panic(err)
	}

	return out.String()
}

func (s bashScript) Shell() ShellName {
	return ShellNameBash
}
