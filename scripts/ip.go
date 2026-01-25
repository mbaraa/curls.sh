package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "ip",
		description: "Returns details about your IP address",
		scriptText: `#!/bin/bash
data=$(curl "https://ipinfo.io" -s)

city=$(echo $data | jq '.city' -r)
regionName=$(echo $data | jq '.region' -r)
country=$(echo $data | jq '.country' -r)
ip=$(echo $data | jq '.ip' -r)
echo "[$ip]" $city, $regionName in $country.`,
	})
}
