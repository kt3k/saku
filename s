#!/bin/sh
set -eu
PJ=saku

# get the version tag from $PJ.md
VERSION=$(sed -e 's/^.*'$PJ'  *v*\([0-9]*\.[0-9]*\.[0-9]\).*$/\1/p;d' < $PJ.md)
TAG=v$VERSION
VERSION_DIR=$HOME/.$PJ/$TAG
EXE=$VERSION_DIR/$PJ

# Download it if the binary doesn't exist.
if [ ! -e "$EXE" ]; then
  echo "Downloading the $PJ binary (version $VERSION)"
  TEMP=$(mktemp -d)
  curl -fsSL https://github.com/kt3k/"$PJ"/releases/download/"$TAG"/${PJ}_"$VERSION"_"$(uname | tr '[:upper:]' '[:lower:]')"_amd64.tar.gz -o "$TEMP"/bin.tgz
  echo "Extracting the $PJ binary"
  mkdir -p "$VERSION_DIR"
  # gunzip ...
  tar -C "$VERSION_DIR" -xf "$TEMP"/bin.tgz "$PJ"
  rm -rf "$TEMP"
  chmod +x "$EXE"
fi

echo "Use $EXE"

# Run $PJ
$EXE "$@"
