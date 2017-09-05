#!/bin/bash - 
#===============================================================================
#
#          FILE: loop.sh
# 
#         USAGE: ./loop.sh 
# 
#   DESCRIPTION: 
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: YOUR NAME (), 
#  ORGANIZATION: 
#       CREATED: 09/05/2017 14:01
#      REVISION:  ---
#===============================================================================

set -o nounset                              # Treat unset variables as an error
for i in {1..99}; 
do 
    if [ "$(( $i % 2 ))" != "0" ]
    then
        echo $i; 
        fi
    done

