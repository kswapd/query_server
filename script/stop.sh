#!/bin/sh
pid=`ps -ef |grep query_server | grep -v grep | awk '{print $2}'`
kill -9 $pid