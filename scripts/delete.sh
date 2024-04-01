#! /bin/sh

curl --request DELETE -sL \
     --url 'localhost:9999/engine'\
     --header 'Content-Type: application/json' \
     --data '[1774757765512695808]'