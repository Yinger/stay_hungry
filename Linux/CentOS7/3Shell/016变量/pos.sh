#!/bin/bash

# $1 $2...$9 ${10}

# # # pos1=$1
# # # pos2=$2

# # # echo $pos1
# # # echo $pos2

# # # 省略形式
# # echo $1
# # echo $2

# # 规避空值
# pos1=$1
# pos2=${2}_

# echo $1
# echo $2

# 参数替换
pos1=$1
pos2=${2-_}

echo $1
echo $2