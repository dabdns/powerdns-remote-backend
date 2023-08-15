#!/bin/bash

cd "$(mktemp -d)"
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')"
KREW="krew-${OS}_${ARCH}"
curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz"
tar zxvf "${KREW}.tar.gz"
./"${KREW}" install krew

if [ -f "$HOME/.zshrc" ]; then
cat << EOF >> $HOME/.zshrc
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
EOF
fi

if [ -f "$HOME/.bashrc" ]; then
cat << EOF >> $HOME/.bashrc
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
EOF
fi
