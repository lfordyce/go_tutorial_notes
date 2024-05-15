#!/bin/bash
#
host="localhost"
port=8080
echo -n "PING" >/dev/udp/$host/$port
