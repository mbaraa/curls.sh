package scripts

import (
	"strings"
)

// ShellName represents name of a used shell in a script.
type ShellName string

const (
	ShellNameBash ShellName = "bash"
)

// Script represent a script template.
type Script interface {
	// Id returns the script's id.
	Id() string
	// Description returns description about the script.
	Description() string
	// RunString accepts a slice of strings [args] that represents ordered shell script arguments
	RunString(args []string) string
	// Shell returns name of the used shell with the script.
	Shell() ShellName
}

func allScripts() []Script {
	outScripts := make([]Script, 0, len(bashScripts))
	for _, script := range bashScripts {
		outScripts = append(outScripts, script)
	}

	return outScripts
}

// All returns a slice of available scripts.
func All() []Script {
	return allScripts()
}

// AllByShell returns the available scripts sorted by shells.
func AllByShell() map[ShellName][]Script {
	sbs := make(map[ShellName][]Script)

	sbs[ShellNameBash] = make([]Script, len(bashScripts))
	for _, script := range allScripts() {
		sbs[ShellNameBash] = append(sbs[ShellNameBash], script)
	}

	return sbs
}

// AllById returns the available scripts sorted by their IDs.
func AllById() map[string]Script {
	sbn := make(map[string]Script)

	for _, script := range allScripts() {
		sbn[script.Id()] = script
	}

	return sbn
}

func idify(content string) string {
	cleanContent := strings.ToLower(content)

	var idBuilder strings.Builder

	for _, ch := range cleanContent {
		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == ' ' || ch == '_' {
			idBuilder.WriteRune(ch)
		} else {
			idBuilder.WriteRune('_')
		}
	}

	slug := strings.ReplaceAll(idBuilder.String(), " ", "_")
	for strings.Contains(slug, "__") {
		slug = strings.ReplaceAll(slug, "__", "_")
	}

	slug = strings.Trim(slug, "_")

	return slug
}
