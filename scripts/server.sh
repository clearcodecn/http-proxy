#!/bin/bash

pidFile=/tmp/server.pid

function start(){
  nohup server > /tmp/proxy.log 2>&1 &
  pid=$!
  echo $pid > $pidFile
}

function stop() {
  if [[ -f $pidFile ]] ; then
    kill -9 $(cat $pidFile)
    rm -rf $pidFile
  fi
}

case $1 in
"start")
    start
    ;;
"stop")
    stop
    ;;
"reload")
    stop
    start
    ;;
esac

exit 0
