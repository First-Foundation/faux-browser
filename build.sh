#!/usr/bin/env bash

# Based on https://gist.github.com/eduncan911/68775dba9d3c028181e4

type setopt >/dev/null 2>&1

contains() {
    # Source: https://stackoverflow.com/a/8063398/7361270
    [[ $1 =~ (^|[[:space:]])$2($|[[:space:]]) ]]
}

# Delete existing binaries
rm -r ./bin

# Create binary directories
mkdir ./bin

# Post-build function
post_build () {
    echo "--- Performing post-build actions on $1"
    upx -9 $1
}

# Build faux-browser for all targets
FAILURES=""
REMOVEDS=""

echo "Building all faux-browsers..."

while IFS= read -r target; do
    GOOS=${target%/*}
    GOARCH=${target#*/}
    BIN_FILENAME="./bin/faux-browser-${GOOS}-${GOARCH}"
    if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
    CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags \"-s -w\" -trimpath -o ${BIN_FILENAME} ./faux-browser.go"
    echo "--- Building ${BIN_FILENAME}"
    eval "${CMD}" || FAILURES="${FAILURES} ${GOOS}/${GOARCH}"
    post_build "${BIN_FILENAME}"
    echo ""
done <<< "$(go tool dist list)"

if [[ "${FAILURES}" != "" ]]; then
    echo ""
    echo "faux-browser build failed on: ${FAILURES}"
    exit 1
fi

