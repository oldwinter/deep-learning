#docker image version
VERSION=$1  # eg 0222-1

#huaweicloud north region docker registry and organization 
SWR_ADDR=$2 # eg 100.125.0.198:20202/istio-new-demo

docker tag "new-demo/frontend:v1-${VERSION}"        "${SWR_ADDR}/frontend:v1-${VERSION}"
docker tag "new-demo/frontend:v2-${VERSION}"        "${SWR_ADDR}/frontend:v2-${VERSION}"
docker tag "new-demo/forecast:v1-${VERSION}"        "${SWR_ADDR}/forecast:v1-${VERSION}"
docker tag "new-demo/forecast:v2-${VERSION}"        "${SWR_ADDR}/forecast:v2-${VERSION}"
docker tag "new-demo/advertisement:v1-${VERSION}"   "${SWR_ADDR}/advertisement:v1-${VERSION}"
docker tag "new-demo/recommendation:v1-${VERSION}"  "${SWR_ADDR}/recommendation:v1-${VERSION}"

docker push "${SWR_ADDR}/frontend:v1-${VERSION}"
docker push "${SWR_ADDR}/forecast:v1-${VERSION}"
docker push "${SWR_ADDR}/forecast:v2-${VERSION}"
docker push "${SWR_ADDR}/advertisement:v1-${VERSION}"
docker push "${SWR_ADDR}/recommendation:v1-${VERSION}"