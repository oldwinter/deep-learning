#docker image version
VERSION=$1  # 0222-1

#gamma huaweicloud north region docker registry and organization 
SWR_ADDR=$2 # 100.125.5.235:20202/istio-new-demo

##huaweicloud north region docker registry and organization 
PRE_FIX=$3 #swr.cn-north-1.myhuaweicloud.com/istio-new-demo

docker tag "$AA/frontend:v1-${VERSION}"        "${SWR_ADDR}/frontend:v1-${VERSION}"
docker tag "$AA/forecast:v1-${VERSION}"        "${SWR_ADDR}/forecast:v1-${VERSION}"
docker tag "$AA/forecast:v2-${VERSION}"        "${SWR_ADDR}/forecast:v2-${VERSION}"
docker tag "$AA/advertisement:v1-${VERSION}"   "${SWR_ADDR}/advertisement:v1-${VERSION}"
docker tag "$AA/recommendation:v1-${VERSION}"  "${SWR_ADDR}/recommendation:v1-${VERSION}"

docker push "${SWR_ADDR}/frontend:v1-${VERSION}"
docker push "${SWR_ADDR}/forecast:v1-${VERSION}"
docker push "${SWR_ADDR}/forecast:v2-${VERSION}"
docker push "${SWR_ADDR}/advertisement:v1-${VERSION}"
docker push "${SWR_ADDR}/recommendation:v1-${VERSION}"
