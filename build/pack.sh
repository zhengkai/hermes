#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$(dirname "$DIR")/dist/release"

if [ ! -f "hermes-next" ]; then
	exit 1
fi

TARGET="hermes_${1}.tgz"

mv hermes-next hermes

if [ -f "$TARGET" ]; then
	rm "$TARGET" || :
fi

tar zcvf "$TARGET" hermes
