#!/bin/sh

while true
do

    CHECK=$(pgrep kubectl | grep "" -c)
    if [ $CHECK -eq 0 ]
    then
        kubectl port-forward svc/prometheus-server 9090:80
    fi 
    sleep 3

done