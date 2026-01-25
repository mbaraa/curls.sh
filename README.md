# curls.sh

A bunch of completely safe and cool scripts to run using the classic `curl thing.com | sh`.

PS: the scripts were only tested on Linux and FreeBSD, no idea what's gonna happen on other systems, obviously they won't work on Win\*ows...

## Usage

Pretty straight forward, just run

```bash
curl curls.sh/script | sh
```

To provide CLI arguments before hand, pass them as `args` query parameter

```bash
curl curls.sh/script?args="-pr" | sh
```

## Available Scripts

1. `tbw`: `curl curls.sh/tbw?args="/dev/nvme0n1" | bash`; Reports Tera Bytes Written of the provided SSD device.
   - requires [smartmontools](https://pkgs.org/search/?q=smartmontools) to be installed.
2. `samsung_bloat_remover`: `curl curls.sh/samsung_bloat_remover | bash`; Removes Samsung's One UI pre installed bloat.
   - requires [android-tools](https://developer.android.com/tools/releases/platform-tools) to be installed.
3. `init_debian_server`: `curl curls.sh/init_debian_server | bash`; Installs Docker and nginx (with certbot) on a Debian server.
4. `vbox_macos`: `curl curls.sh/vbox_macos | bash`; Fixes macOS virtual box machine.
   - requires [VirtualBox](https://virtualbox.org) :trollface:
5. `ulfius_fedora`: `curl curls.sh/ulfius_fedora | bash`; Installs C build tools and builds and installs [Ulfius](https://github.com/babelouest/ulfius) for Fedora.
6. `random_emoji`: `curl curls.sh/random_emoji | bash`; prints a random emoji every second.
7. `ip`: `curl curls.sh/ip | bash`; prints details about your IP using [ipinfo.io]
   - requires [jq]()

## Contributing

I bet you have something else interesting to add to the scripts list and it would be really nice of you to contribute, check the poorly written [CONTRIBUTING.md](/CONTRIBUTING.md) for more info.
