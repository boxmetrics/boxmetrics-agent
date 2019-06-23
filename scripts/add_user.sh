#! /bin/bash

declare -a groups

while getopts u:p:g: option; do
	case "${option}" in

	u) username=${OPTARG} ;;
	p) password=${OPTARG} ;;
	g) groups+=(${OPTARG}) ;;
	esac
done

group=''

for i in ${groups[@]}; do
	if [ $i = ${groups[0]} ]; then
		group="$i"
	else
		group+=",$i"
	fi
done

# Script to add a user to Linux system
if [ $(id -u) -eq 0 ]; then
	egrep "^$username" /etc/passwd >/dev/null
	if [ $? -eq 0 ]; then
		echo "$username exists!"
		exit 1
	else
		pass=$(perl -e 'print crypt($ARGV[0], "password")' $password)
		if [ $group != '' ]; then
			useradd -m -U -G $group -p $pass $username
		else
			useradd -m -U -p $pass $username
		fi
		[ $? -eq 0 ] && echo "User has been added to system!" || echo "Failed to add a user!"
	fi
else
	echo "Only root may add a user to the system"
	exit 2
fi
