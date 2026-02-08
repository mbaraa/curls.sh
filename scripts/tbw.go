package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "tbw",
		description: "Displays SSD's Tera bytes written",
		scriptText: `#!/bin/bash

SSD_DEVICE="{{ .V1 }}"

if [[ "$SSD_DEVICE" == "" ]]; then
    SSD_DEVICE="/dev/sda"
fi

ON_TIME_TAG="Power On Hours"
WEAR_COUNT_TAG="Available Spare"
LBAS_WRITTEN_TAG="Data Units Written"
LBA_SIZE=512000 # Value in bytes

BYTES_PER_MB=1048576
BYTES_PER_GB=1073741824
BYTES_PER_TB=1099511627776

SMART_INFO=$(sudo /usr/bin/smartctl -A "$SSD_DEVICE")

# Extract required attributes
ON_TIME=$(echo "$SMART_INFO" | grep "$ON_TIME_TAG" | awk '{print $4}' | sed 's/,//g')
WEAR_COUNT=$(echo "$SMART_INFO" | grep "$WEAR_COUNT_TAG" | awk '{print $3}' | sed 's/Threshold://' | sed 's/,//g')
LBAS_WRITTEN=$(echo "$SMART_INFO" | grep "$LBAS_WRITTEN_TAG" | awk '{print $4}' | sed 's/,//g')

# Convert LBAs -> bytes
BYTES_WRITTEN=$(echo "$LBAS_WRITTEN * $LBA_SIZE" | bc)
MB_WRITTEN=$(echo "scale=3; $BYTES_WRITTEN / $BYTES_PER_MB" | bc)
GB_WRITTEN=$(echo "scale=3; $BYTES_WRITTEN / $BYTES_PER_GB" | bc)
TB_WRITTEN=$(echo "scale=3; $BYTES_WRITTEN / $BYTES_PER_TB" | bc)

# Tada
echo "------------------------------"
echo " SSD Status:   $SSD_DEVICE"
echo "------------------------------"
echo " On time:      $(echo $ON_TIME | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta') hr"
echo "------------------------------"
echo " Data written:"
echo "           MB: $(echo $MB_WRITTEN | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta')"
echo "           GB: $(echo $GB_WRITTEN | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta')"
echo "           TB: $(echo $TB_WRITTEN | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta')"
echo "------------------------------"
echo " Mean write rate:"
echo "        MB/hr: $(echo "scale=3; $MB_WRITTEN / $ON_TIME" | bc | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta')"
echo "------------------------------"
echo " Drive health: ${WEAR_COUNT}"
echo "------------------------------"`,
	})
}
