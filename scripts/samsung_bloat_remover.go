package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "samsung_bloat_remover",
		description: "Removes Samsung One UI's bloatware",
		scriptText: `#!/bin/bash

for i in com.netflix.mediaclient \
    com.spotify.music \
    com.google.android.apps.youtube.music \
    com.facebook.katana \
    com.facebook.services \
    com.facebook.system \
    com.facebook.appmanager \
    com.google.mainline.adservices \
    com.google.android.apps.docs.editors.sheets \
    com.google.android.adservices.api \
    com.google.android.apps.messaging \
    com.google.android.apps.docs \
    com.google.android.apps.photos \
    com.google.android.apps.tachyon \
    com.google.android.feedback \
    com.google.android.googlequicksearchbox \
    com.google.android.inputmethod.latin \
    com.google.android.marvin.talkback \
    com.google.ar.core \
    com.samsung.android.app.settings.bixby \
    com.samsung.android.bixby.wakeup \
    com.samsung.android.bixby.agent \
    com.samsung.android.bixbyvision.framework \
    com.sec.android.app.shealth \
    com.samsung.android.arzone \
    com.samsung.android.tvplus \
    com.samsung.android.app.watchmanagerstub \
    com.samsung.android.app.watchmanager \
    com.samsung.android.waterplugin \
    com.samsung.android.accessibility.talkback \
    com.sec.android.app.sbrowser \
    com.sec.android.easyMover.Agent \
    com.samsung.android.voc \
    com.samsung.android.scloud \
    com.samsung.android.service.aircommand \
    com.samsung.android.samsungpassautofill \
    com.samsung.android.samsungpass \
    com.samsung.android.spayfw \
    com.samsung.android.kidsinstaller \
    com.samsung.android.app.camera.sticker.facearavatar.preload \
    com.samsung.android.globalpostprocmgr \
    com.samsung.android.app.find \
    com.samsung.android.smartswitchassistant \
    com.samsung.android.game.gametools \
    com.samsung.android.game.gos \
    com.samsung.android.game.gamehome \
    com.samsung.android.bixby.ondevice.engb \
    com.samsung.android.bixby.ondevice.arae \
    com.samsung.android.smartface.overlay \
    com.samsung.android.smartcallprovider \
    com.samsung.android.smartface \
    com.samsung.android.smartsuggestions \
    com.sec.android.smartfpsadjuster \
    com.sec.android.mimage.avatarstickers \
    com.android.avatarpicker \
    com.google.android.apps.bard \
    com.microsoft.office.outlook \
    com.microsoft.office.officehubrow \
    com.microsoft.skydrive \
    com.microsoft.appmanager \
    com.linkedin.android \
    com.google.mainline.telemetry \
    com.google.android.videos \
    ; do
    echo "uninstalling $i";
    pm uninstall --user 0 $i
done`,
	})
}
