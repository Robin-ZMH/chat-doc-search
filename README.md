# Start the API server

## How to run locally

1. Clone the repository
2. Run the following command to start the server:
```bash
make start
```

## How to check logs

Run the following command to check the logs:
```bash
make logs
```

## How to stop the local service

Run the following command to stop the server:
```bash
make stop
```

# API documentation

## Query(find-many) chat docs

POST /engine/query  

Request Body: User's input prompt(string)

Response Body:
```json
[
  {
    "id": 1775002557081456640,
    "prompt": "string",
    "response": "string"
  }
]
```

example:
```bash
curl --request POST -sL \
    --url 'localhost:9999/engine/query'\
    --header 'Content-Type: application/json' \
    --data '"what is jarvis bot"'
# get
# [{"id":3,"prompt":"What is your sex ?","response":"I am a bot ."},{"id":1,"prompt":"What is your name ?","response":"My name is jarvis ."}]
```

## Insert(in batch) chat docs

POST /engine/insert

Request Body: 
```json
[
  {
  "prompt": "string",
  "response": "string"
  }
]
```

example:
```bash
curl --request POST -sL \
     --url 'localhost:9999/engine/insert'\
     --header 'Content-Type: application/json' \
     --data '[{"prompt":"Is Golang object oriented?","response":"It can be."}]'
```

## Update(in batch) chat docs

PUT/PATCH /engine

Request Body: 
```json
[
  {
    "id": 1775002557081456640,
    "prompt": "string",
    "response": "string"
  }
]
```

example:
```bash
curl --request PUT -sL \
     --url 'localhost:9999/engine'\
     --header 'Content-Type: application/json' \
     --data '[{"prompt":"Is Golang object oriented?","response":"Well, I think so~"}]'
```

## Delete(in batch) chat docs
DELETE /engine

example:
```bash
curl --request DELETE -sL \
     --url 'localhost:9999/engine'\
     --header 'Content-Type: application/json' \
     --data '[1774757765512695808,1775004076325474304]'
```