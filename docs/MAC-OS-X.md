# `MAC-OS-X`

## A note on Mac OSX

I use Linux mainly and the GNU versions of various utilities when on Mac OSX. To install the GNU tools, run the commands below:

```sh
mkdir brew-logs
cd brew-logs
brew install coreutils findutils gnu-tar gnu-sed gawk gnutls gnu-indent gnu-getopt grep make 2>&1 | tee install-log-gnu-tools-install.log
brew upgrade coreutils findutils gnu-tar gnu-sed gawk gnutls gnu-indent gnu-getopt grep make 2>&1 | tee install-log-gnu-tools-upgrade.log
brew node yarn 2>&1 | tee install-log-node-install.log
brew install yarn 2>&1 | tee install-log-yarn-install.log
brew reinstall yarn 2>&1 | tee install-log-yarn-reinstall.log
brew reinstall node 2>&1 | tee install-log-node-reinstall.log
```

Restart your shell then:

```sh
# List all the install binaries and their versions
$ ls -1 /usr/local/opt/*/libexec/gnubin/* | xargs -IV sh -c 'V --version'
...
[ (GNU coreutils) 8.32
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Kevin Braunsdorf and Matthew Bradburn.
b2sum (GNU coreutils) 8.32
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
...
```

Credits to <https://gist.github.com/abouteiller> for this post <https://gist.github.com/skyzyx/3438280b18e4f7c490db8a2a2ca0b9da#gistcomment-3049694> for the below snippet. Once the above commands have completed add the below to `~/.zshrc`, or if you prefer you can see part of my `~/.zshrc` in [`./scripts/.zshrc_brew`](./scripts/.zshrc_brew):

```sh
if type brew &>/dev/null; then
  HOMEBREW_PREFIX=$(brew --prefix)
  # gnubin; gnuman
  for d in ${HOMEBREW_PREFIX}/opt/*/libexec/gnubin; do export PATH=$d:$PATH; done
  # I actually like that man grep gives the BSD grep man page
  #for d in ${HOMEBREW_PREFIX}/opt/*/libexec/gnuman; do export MANPATH=$d:$MANPATH; done
fi
```

After restarting shell:

```sh
$ make --version
GNU Make 4.3
Built for x86_64-apple-darwin19.2.0
Copyright (C) 1988-2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

## Go Install

### Mac OSX

[`./scripts/install-go.sh`](./scripts/install-go.sh)
