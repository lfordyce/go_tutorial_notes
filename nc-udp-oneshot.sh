#!/bin/sh

helpword=$(nc -h 2>&1 | awk '{print$1;exit}')

case $helpword in
  *GNU*) args=-uc ;;
  *) args=-uq0 ;;
esac
exec nc $args "$@"

