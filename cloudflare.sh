#!/bin/bash
curl https://api.filebase.io/v1/ipfs/pins -H "Authorization: Bearer NENDRUJFREZBMzYxMDZCQzZCNzk6V0FYZGM0ZkhEOUxDcHQ3VE5HdlZVZktBa0RRdUpmWVA4b3RCMnZtbjpqb2JjaGVy" > filebase.txt
DNSTXT=$(cat filebase.txt | sed 's/,/\n/g' | grep "cid" | sed 's/:/\n/g' | sed -n '3p' |sed 's/"//g')
DNSLINK="dnslink=$DNSTXT"

# 修改DNS ID
curl -s -X PUT "https://api.cloudflare.com/client/v4/zones/5589d7d645b0fc8c0c4be50bb19b0112/dns_records/67ede0849e78191be3fbd0793e7d789c" \
     -H "Content-Type:application/json" \
     -H "X-Auth-Key:4da7be8506a9f9649e0901de3e7db9a066ff8" \
     -H "X-Auth-Email:13028911306@163.com" \
     --data '{"type":"txt","name":"_dnslink.ipfs.jobcher.com","content":"'$DNSLINK'","ttl":1,"proxied":false}'