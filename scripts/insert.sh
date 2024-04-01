#! /bin/sh

curl --request POST -sL \
     --url 'localhost:9999/engine/insert'\
     --header 'Content-Type: application/json' \
     --data '[{"prompt":"what is a circle","response":"the earth is circle"}]'

