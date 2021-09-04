#!/bin/bash

if [ $# -eq 0 ]
then
	echo
	echo "Usage : $0 <SRCDIR>"
	exit 0
fi

# Check if the user is root
if [ "$UID" != "0" ]
then
	echo "[-] The script needs root user"
	exit 1
fi;

# Return code
RETURN_CODE=0

# Get the directory where to get the script and the configuration
SRCDIR=$*

# Define the files to copy
SSH_SCRIPT="$SRCDIR/ssh-notificate.sh"
PRM_FILE="$SRCDIR/webhook.prm"

# Define the arrival directory
DST_DIR="/etc/ssh"

# Define a function to copy a file and set the permission
function copy_file(){
	local PERMISSION="$1"
	local SRC_PATH="$2"
	local DST_PATH="$3"

	cp "$SRC_PATH" "$DST_PATH"
	RETURN_CODE=$([ $? == 0 ] && echo "$RETURN_CODE" || echo "1")

	chown root:root "$DST_PATH"
	RETURN_CODE=$([ $? == 0 ] && echo "$RETURN_CODE" || echo "1")

	chmod "$PERMISSION" "$DST_PATH"
	RETURN_CODE=$([ $? == 0 ] && echo "$RETURN_CODE" || echo "1")
}

# Copy the script and the configuration
echo "[+] Installing sshrc script with configuration"
copy_file "755" "$SSH_SCRIPT" "$DST_DIR/sshrc"
copy_file "644" "$PRM_FILE" "$DST_DIR/webhook.prm"

exit "$RETURN_CODE"
