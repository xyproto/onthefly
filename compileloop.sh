#!/bin/sh
SOURCE=main.go
BIN=main
PIDFILE=$BIN.pid
LOG=error.log
M5=nop
SUMFILE=sumfile.txt
echo 'Starting compilation loop'
echo 'Reading pid'
if [ -e $PIDFILE ]; then
  echo 'Killing server'
  kill `cat $PIDFILE` > /dev/null
  rm $PIDFILE
fi
while true; do
  OLDM5=$M5
  md5sum *.go > $SUMFILE
  M5=$(md5sum $SUMFILE)
  if [ "$OLDM5" != "$M5" ]; then
    echo 'Source changed'
    echo 'Reading pid'
    if [ -e $PIDFILE ]; then
      echo 'Killing server'
      kill `cat $PIDFILE` > /dev/null
      rm $PIDFILE
    fi
    clear
    date
    echo
    echo -n 'Recompiling...'
    [ -e $LOG ] && rm $LOG
    go build -o $BIN > $LOG
    if [ "$(wc -c $LOG | cut -d' ' -f1)" == '0' ]; then
      rm $LOG
    fi
    if [ -e $LOG ]; then
      echo
      cat $LOG
    else
      echo ok
    fi
    echo
    echo 'Starting server'
    ./$BIN &
    echo 'Writing pid'
    pgrep $BIN > $PIDFILE
  fi
  # Wait for the source to be changed
  inotifywait *.go
done
