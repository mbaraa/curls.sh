package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "ulfius_fedora",
		description: "Installs Ulfius for Fedora",
		scriptText: `#!/bin/bash

init() {
    sudo dnf install cmake gcc pkg-config libmicrohttpd-devel jansson-devel libcurl-devel gnutls-devel libgcrypt-devel zlib-devel git
}

install_orcania() {
    git clone https://github.com/babelouest/orcania.git
    cd orcania/src
    make -j$nproc DESTDIR=/usr
    sudo make install DESTDIR=/usr
    sudo make install
    cd ../../
}

install_yder() {
    git clone https://github.com/babelouest/yder.git
    cd yder/src
    make -j$nproc DESTDIR=/usr
    sudo make install DESTDIR=/usr
    cd ../../
}

install_ulfius() {
    git clone https://github.com/babelouest/ulfius
    cd ulfius/src
    make -j$nproc WEBSOCKETFLAG=1 DESTDIR=/usr
    sudo make install DESTDIR=/usr
    cd ../../
}

cleanup() {
    rm -rf orcania yder ulfius
}

init
install_orcania
install_yder
install_ulfius
#cleanup`,
	})
}
