# Go

## Install Go

```console
export GO_VERSION=1.13.6
curl -sfLo go$GO_VERSION.linux-amd64.tar.gz https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

nano ~/.profile
--add the following line to the bottom of the file
export PATH=$PATH:/usr/local/go/bin
-- save and exit
source $HOME/.profile
go version
```

## Install `direnv`
```console
export DE_VERSION=2.20.0
curl -sfLo direnv https://github.com/direnv/direnv/releases/download/v$DE_VERSION/direnv.linux-amd64
chmod +x direnv
sudo mv direnv /usr/local/bin

nano ~/.bashrc
--add the following line to the bottom of the file
eval "$(direnv hook bash)"
-- save and exit
exec bash
```