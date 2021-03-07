#!/bin/bash

pidFile=/tmp/client.pid

function start(){
  nohup client > /tmp/proxy.log 2>&1 &
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
    start()
    ;;
"stop")
    stop()
    ;;
"reload")
    stop()
    start()
esac

exit 0