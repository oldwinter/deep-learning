#!/bin/bash

VERSION="0222-1"

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )
SCRIPTDIR="$DIR/service"

pushd "$SCRIPTDIR/frontend-reactjs"
    docker build -t "new-demo/frontend:${VERSION}" .
popd

pushd "$SCRIPTDIR/forecast-nodejs"
    docker build -t "new-demo/forecast:${VERSION}" .
popd

pushd "$SCRIPTDIR/advertisement-golang"
    docker build -t "new-demo/advertisement:${VERSION}" .
popd

pushd "$SCRIPTDIR/news-golang"
    docker build -t "new-demo/news:${VERSION}" .
popd

pushd "$SCRIPTDIR/recommondation-redis"
    docker build -t "new-demo/recommondation:${VERSION}" .
popd

