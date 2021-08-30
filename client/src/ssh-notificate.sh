#!/bin/bash

# Define the time when the script is called and so when the user connects
CONNECT_TIME=$(date '+%Y-%m-%d %H:%M:%S')

# Define the parameter file path
DIRNAME=$(dirname "$(readlink -f "$0")")
PRM_FILE="$DIRNAME/webhook.prm"

# Define the useful parameters
HOSTNAME=$(hostname)

IP_SRC=$(echo "$SSH_CONNECTION"   | awk -F' ' '{print $1}')
PORT_SRC=$(echo "$SSH_CONNECTION" | awk -F' ' '{print $2}')
IP_DST=$(echo "$SSH_CONNECTION"   | awk -F' ' '{print $3}')
PORT_DST=$(echo "$SSH_CONNECTION" | awk -F' ' '{print $4}')

DNS_SRC_REVERSE=$(dig -x "$IP_SRC" +short 2>/dev/null)
if [ "$DNS_SRC_REVERSE" != "" ]
then
	SRC_MSG=" ($DNS_SRC_REVERSE)"
fi

# Log the connection
MESSAGE="User '$USER' has just logged in from $IP_SRC$SRC_MSG with port $PORT_SRC to $IP_DST ($HOSTNAME) with port $PORT_DST at $CONNECT_TIME"
echo "[i] $MESSAGE"
logger -t ssh-wrapper "[i] $MESSAGE"

# Push the notification through the webhook
if [ -f "$PRM_FILE" ]
then
	SSHCONNECT_WEBHOOK_URL=$(grep "^https://" "$PRM_FILE" | tail -n1)
	if [ "$SSHCONNECT_WEBHOOK_URL" = "" ]
	then
		logger -t ssh-wrapper "[-] Error no webhook URL defined"
	else
		CURL_COUNT=0
		while [ "$CURL_COUNT" -le 3 ]
		do
			CURL_COUNT=$(( CURL_COUNT+1 ))
			curl -X POST -H "Content-Type: application/json"  -d "{\"content\":\"$MESSAGE\"}" "$SSHCONNECT_WEBHOOK_URL" 1>/dev/null 2>&1
			RSLT="$?"
			if [ "$RSLT" -eq 0 ]
			then
				break
			fi
		done

		if [ "$RSLT" -ne 0 ]
		then
			logger -t ssh-wrapper "[-] Error in the curl command to the webhook to notify an SSH connection [error code ($RSLT)]"
		else
			logger -t ssh-wrapper "[i] Notification sent after $CURL_COUNT attempt(s)"
		fi
	fi
else
	echo "[-] No parameter file found"
	logger -t ssh-wrapper "[-] Error no parameter file found to notify an SSH connection"
fi

