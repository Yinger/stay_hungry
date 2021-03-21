#!/bin/bash

# demo if then if then fi fi

if [ $UID = 0 ]; then
    echo "please run"
    if [ -x /tmp/elif.sh ] ; then
        /tmp/elif.sh
    fi
else 
    echo "switch user root"
fi