set -o errexit

if [ "$#" -ne 1 ]; then
    echo Missing version parameter
    echo Usage: build-services.sh \<version\>
    exit 1
fi

#docker image version
VERSION=$1 # eg: 0222-1

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )
SCRIPTDIR="$DIR/service"

pushd "$SCRIPTDIR/frontend-reactjs"
  pushd v1
    docker build -t "new-demo/frontend:v1-${VERSION}" .
  popd
  pushd v1
    docker build -t "new-demo/frontend:v1-${VERSION}" .
  popd
popd

pushd "$SCRIPTDIR/forecast-nodejs"
  pushd v1
    docker build -t "new-demo/forecast:v1-${VERSION}" .
  popd
  pushd v2
    docker build -t "new-demo/forecast:v2-${VERSION}" .
  popd
popd

pushd "$SCRIPTDIR/advertisement-golang"
  pushd v1
    docker build -t "new-demo/advertisement:v1-${VERSION}" .
  popd
  pushd v2
    docker build -t "new-demo/advertisement:v2-${VERSION}" .
  popd
popd

pushd "$SCRIPTDIR/recommendation-java"
  pushd v1
    docker build -t "new-demo/recommendation:v1-${VERSION}" .
  popd
  pushd v2
    docker build -t "new-demo/recommendation:v2-${VERSION}" .
  popd
popd

pushd "$SCRIPTDIR/news-golang"
  pushd v1
    docker build -t "new-demo/news:v1-${VERSION}" .
  popd
popd

