#!/bin/bash

PUBLIC_IP_QUERY_HOST="${PUBLIC_IP_QUERY_HOST:-icanhazip.com}"

if [[ -z $CANARY_FQDN ]]; then
    echo "CANARY_FQDN not found in environment!"
    exit 1
fi

PUBLIC_IP_ADDRESS=`wget -qO - ${PUBLIC_IP_QUERY_HOST}`
DNS_IP_ADDRESS=`dig +short ${CANARY_FQDN}`

if [ "$PUBLIC_IP_ADDRESS" = "$DNS_IP_ADDRESS" ]
then ADDRESS_MATCH=true
else ADDRESS_MATCH=false
fi

printf '{"canary": {"match": %s, "public_ip_address": "%s", "dns_ip_address": "%s"}}\n' "${ADDRESS_MATCH}" "${PUBLIC_IP_ADDRESS}" "${DNS_IP_ADDRESS}"
