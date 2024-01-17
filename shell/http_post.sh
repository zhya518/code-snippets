#!/bin/bash

cat audit.txt | while read mobile
do
#mobile="86912087831"
body="{\"sequence\":1662109887947,\"infos\":[{\"mobile\":\"${mobile}\",\"app_id\":\"ikxd\"}],\"status\":0}"
echo $body

curl 'https://boss-proxy.ihago.cn/boss_proxy/ymicro/api?group_id=881' \
  -H 'content-type: application/json' \
  -H 'x-ymicro-api-method-name: Uaasadmin.AuditAccount' \
  -H 'x-ymicro-api-service-name: net.ihago.ymicro.srv.uaasadmin' \
  --data-raw ${body} \
  --compressed

done
