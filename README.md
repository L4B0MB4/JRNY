# JRNY (Journey)

[![Go Reference](https://pkg.go.dev/badge/github.com/L4B0MB4/JRNY.svg)](https://pkg.go.dev/github.com/L4B0MB4/JRNY)

Captures and builds a traceable and editable journey of an object (e.g. customer) through time without any cloud tools needed.

(This is for now a project to improve my own `GO` knowledge)

# Docker

## RabbitMQ

Locally currently using

```bash
docker run --hostname=my-rabbit -p 8080:15672 -p 5672:5672 -d rabbitmq:3-management
```

for development

# Requests

## Example

```bash
curl --location 'http://localhost:8081/api/event' \
--header 'Content-Type: application/json' \
--data '{
    "type":"abc",
    "id":"ed4c5e8f-c512-48ba-b488-bb4be07508e3",
    "attributes":{
        "hallo":"h"
    },
    "relationships":{
        "cde":[{
            "type":"cde",
            "id":"5d4c5e8f-c512-48ba-b488-bb4be07508e3"
        }]
    }
}'
```
