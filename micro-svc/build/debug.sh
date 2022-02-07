#!/bin/bash

# 若简单部署在docker中，则使用一个新的网桥
docker network create test_net


VERSION="0222-1" 

docker run \
-p 3000:3000 \
-e SVC_FORECAST="http://svc-forecast:3002" \
-e SVC_ADVERTISEMENT="http://svc-advertisement:3003" \
--network test_net \
--network-alias svc-frontend \
new-demo/frontend:2020-1

docker run \
-p 3002:3002 \
--network test_net \
--network-alias svc-forecast \
new-demo/forecast:2020-1


docker run \
-p 3003:3003 \
-e NEWS_SERVER_NAME="svc-news:3007" \
--network test_net \
--network-alias svc-advertisement \
new-demo/advertisement:2020-1


docker run  \
-p 3007:3007 \
-e PORT_NUM=":3007" \
--network test_net \
--network-alias svc-news \
new-demo/news:2020-1

docker run \
--network test_net \
--network-alias redis-server \
new-demo/recommondation:2020-1
#redis client
#docker run -it --network test_net --rm redis redis-cli -h redis-server