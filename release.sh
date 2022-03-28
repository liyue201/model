#!/usr/bin/env bash

pushd $(dirname $0) >/dev/null
SCRIPTPATH=$(pwd -P)
popd >/dev/null
SCRIPTFILE=$(basename $0)

function log() {
  echo "================================================================================"
  echo "$(date +'%Y-%m-%d %H:%M:%S%z') [INFO] - $@"
  echo ""
}

VERSION=$(cat ${SCRIPTPATH}/version.txt)

MAJOR=$(echo $VERSION | awk -F'.' '{ print $1 }')
MINOR=$(echo $VERSION | awk -F'.' '{ print $2 }')
PATCH=$(echo $VERSION | awk -F'.' '{ print $3 }')

PATCH=$((${PATCH} + 1))
if [[ ${PATCH} == 10 ]]; then
  PATCH=0
  MINOR=$((${MINOR} + 1))
fi
if [[ ${MINOR} == 10 ]]; then
  MINOR=0
  MAJOR=$((${MAJOR} + 1))
fi

VERSION="${MAJOR}.${MINOR}.${PATCH}"
echo ${VERSION} >${SCRIPTPATH}/version.txt

# git add .
# git commit -m "prepare to publish v${VERSION}"
log "publish overeality-server-model v${VERSION}"
git tag -a "v${VERSION}" -m "publish overeality-server-model v${VERSION}"
git push -u origin v${VERSION}

cd ${SCRIPTPATH}
git add .
git commit -m"[chore] update version.txt for new tag v${VERSION}"
git push
