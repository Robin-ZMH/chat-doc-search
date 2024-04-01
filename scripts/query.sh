#! /bin/sh

curl --request POST -sL \
     --url 'localhost:9999/engine/query'\
     --header 'Content-Type: application/json' \
     --data '"jarvis bot"'