#!/bin/sh
kill -9 pgrep tg-keyword-repl
sleep 10
go build
nohup ./tg-keyword-reply-bot  > admin.log 2>&1 &
