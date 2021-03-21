#!/bin/bash

# help display help help

# for pos in  $*
# do 
#     if [ "$pos" = "help"] ; then
#         echo $pos $pos
#     fi
# done

while [ $# -ge 1 ]
do
    # echo $#
    # echo "do something"
    if [ "$1" = "help" ]; then
        echo $1 $1
    fi

    # shift:参数左移
    # 例：
    # a.sh a b c d
    # 每执行一次 shift 就会删掉一个参数（第一个参数）
    shift
done