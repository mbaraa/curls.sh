package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "random_emoji",
		description: "Prints a rendom emoji every 1s",
		scriptText: `#!/bin/bash
while true; do
	rand=$(shuf -i 2600-2700 -n 1)
	echo -en "   \u$rand"
	sleep 1
done`,
	})
}
