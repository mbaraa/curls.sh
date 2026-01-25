package scripts

func init() {
	bashScripts = append(bashScripts, bashScript{
		id:          "init_debian_server",
		description: "Initialize a Debian server with nginx and Docker",
		scriptText: `#!/bin/bash

basic_setup() {
    sudo apt update
    sudo apt upgrade -y
    sudo apt install vim htop git wget ca-certificates curl lsof neofetch -y

    wget https://github.com/junegunn/fzf/releases/download/v0.62.0/fzf-0.62.0-linux_amd64.tar.gz -O fzf.tar.gz
    tar xzf fzf.tar.gz
    sudo cp ./fzf /usr/bin/fzf

    git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
    yes y | ~/.fzf/install
}

install_docker() {
    for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt remove $pkg -y; done

    sudo install -m 0755 -d /etc/apt/keyrings
    sudo curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
      $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
      sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt update
    sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
    sudo gpasswd -a $USER docker
    sudo systemctl enable docker
}

install_nginx() {
    sudo apt install nginx python3 python3-dev python3-venv libaugeas-dev gcc -y
    sudo python3 -m venv /opt/certbot/
    sudo /opt/certbot/bin/pip install --upgrade pip
    sudo /opt/certbot/bin/pip install certbot certbot-nginx
    sudo ln -s /opt/certbot/bin/certbot /usr/bin/certbot
    sudo systemctl enable nginx
}

basic_setup
install_docker
install_nginx
sudo reboot`,
	})
}
