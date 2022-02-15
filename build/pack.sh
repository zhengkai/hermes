#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$(dirname "$DIR")/dist/release"

if [ -z "$2" ]; then
	exit
fi

GOOS="$1" GOARCH="$2" "${DIR}/build-server.sh" release

PLATFORM="$1"
if [ "$PLATFORM" == "darwin" ]; then
	PLATFORM="mac"
fi
PLATFORM="${PLATFORM}_${2}"

if [ ! -f "hermes-next" ]; then
	exit 1
fi

TARGET="hermes_${PLATFORM}.tgz"

echo "$TARGET"

mv hermes-next hermes

tar zcf "$TARGET" hermes
