package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "vbox_macos",
		description: "Fixs VirtualBox macOS virtual machine",
		scriptText: `#!/bin/bash

VM_NAME="{{ .V1 }}"

if [[ "$VM_NAME" == "" ]]; then
    VM_NAME="macOS"
fi

VBoxManage modifyvm "${VM_NAME}" --cpuidset 00000001 000106e5 00100800 0098e3fd bfebfbff
VBoxManage setextradata "${VM_NAME}" "VBoxInternal/Devices/efi/0/Config/DmiSystemProduct" "iMac14"
VBoxManage setextradata "${VM_NAME}" "VBoxInternal/Devices/efi/0/Config/DmiSystemVersion" "1.0"
VBoxManage setextradata "${VM_NAME}" "VBoxInternal/Devices/efi/0/Config/DmiBoardProduct" "Iloveapple"
VBoxManage setextradata "${VM_NAME}" "VBoxInternal/Devices/smc/0/Config/DeviceKey" "ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc"
VBoxManage setextradata "${VM_NAME}" "VBoxInternal/Devices/smc/0/Config/GetKeyFromRealSMC" 1
VBoxManage setextradata "${VM_NAME}" "VBoxInternal2/EfiGraphicsResolution" "1600x900"`,
	})
}
