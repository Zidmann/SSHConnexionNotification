#!/bin/bash

# Define the time when the script is called and so when the user connects
CONNECT_TIME=$(date '+%Y-%m-%d %H:%M:%S')

# Define the environment file path
DIRNAME=$(dirname "$(dirname "$(readlink -f "$0")")")
ENV_FILE="$DIRNAME/webhook.env"

# Define the useful parameters
HOSTNAME=$(hostname)

IP_SRC=$(echo "$SSH_CONNECTION"   | awk -F' ' '{print $1}')
PORT_SRC=$(echo "$SSH_CONNECTION" | awk -F' ' '{print $2}')
IP_DST=$(echo "$SSH_CONNECTION"   | awk -F' ' '{print $3}')
PORT_DST=$(echo "$SSH_CONNECTION" | awk -F' ' '{print $4}')

DNS_SRC_REVERSE=$(dig -x "$IP_SRC" +short)
if [ "$DNS_SRC_REVERSE" != "" ]
then
	SRC_MSG=" ($DNS_SRC_REVERSE)"
fi

# Log the connection
MESSAGE="$USER($UID) has just logged in from $IP_SRC$SRC_MSG with port $PORT_SRC to $IP_DST ($HOSTNAME) with port $PORT_DST at $CONNECT_TIME"
echo "[i] $MESSAGE"
logger -t ssh-wrapper "$MESSAGE"

# Push the notification through the webhook
if [ -f "$ENV_FILE" ]
then
	source "$ENV_FILE"
	curl -X POST -H "Content-Type: application/json"  -d "{\"username\":\"$HOSTNAME\", \"content\":\"$MESSAGE\"}" "$WEBHOOK_URL"
	RSLT="$?"
	if [ "$RSLT" != "0" ]
	then
		logger -t ssh-wrapper "[-] Error in the curl command to the webhook to notify an SSH connection [error code ($RSLT)]"
	fi
fi

