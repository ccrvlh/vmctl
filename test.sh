#!/bin/bash
if [[ -c /dev/kvm ]]
then
	echo "ok"
else
	echo "notok"
fi
