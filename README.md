# JRNY (Journey)

[![Go Reference](https://pkg.go.dev/badge/github.com/L4B0MB4/JRNY.svg)](https://pkg.go.dev/github.com/L4B0MB4/JRNY)

Captures and builds a traceable and editable journey of an object (e.g. customer) through time without any cloud tools needed.

(This is for now a project to improve my own `GO` knowledge)

## Merging Idea

Merge-Consumers are split up in different areas of responsibility. E.g. 0-10,10-20,20-30 etc (excluding last number).
If in a relationship or in the id of the event itself there is a number that fits into the area of responsibility
the consumer will save it and its relations.
E.g.

```
id: 1
    rel:
        7
        13
```

fits into the areas 0-10 and 10-20. Therfore the both consumers will save the relation data in the following structure:

consumer 0-10 saves it as:

```
1 (has relation to) -> 7,13
7 -> 1,13
```

consumer 10-20

```
13 -> 7,1
```

If a new item arrives with example values

```
id: 2
    rel:
        13
```

It will be added to

consumer 0-10

```
1 (has relation to) -> 7,13
7 -> 1,13
2 -> 13 (new "row")
```

consumer 10-20

```
13 -> 7,1,2 (new "column")
```

when a new column is added this means a merge of data needs to happen.

Lets add another dataset

```
id: 14
    rel:
        2
```

This causes the following structures

consumer 0-10

```
1 (has relation to) -> 7,13
7 -> 1,13
2 -> 1,14 (new "column" ->  a merge needs to happen)
```

consumer 10-20

```
13 -> 7,1,2
14 -> 2 (new "row")
```

Thinking about these small areas this looks like a lot of duplicated data but with the areas splitted up in the space of 128-bit; I expect the load to be way more evenly spread across different consumers and therefore hopefully making the improved speed worth the additional memory

# Docker

## ENV Variables

```
RABBITMQ_URL = e.g. amqp://guest:guest@localhost:5672/
```

## Server

From the root folder you can run

```bash
docker build -t jrny/server -f deployment/docker/server/Dockerfile .
```

to build the server docker file and then

```bash
docker run -it -p 8081:8081 --name jrny_server jrny/server
```

to run it

## Consumer

Similar to server but with adjusted path to consumer dockerfile

## RabbitMQ

Locally currently using

```bash
docker run --name some-rabbit -p 8080:15672 -p 5672:5672 -d --net jrny_net rabbitmq:3-management
```

for development

## Example on running everything

```bash
docker network create jrny_net

docker run --name some-rabbit -p 8080:15672 -p 5672:5672 -d --net jrny_net rabbitmq:3-management

docker run -d -p 8081:8081 --name jrny_server --env RABBITMQ_URL=amqp://guest:guest@some-rabbit:5672/ --net jrny_net jrny/server

docker run -d --name jrny_consumer --env RABBITMQ_URL=amqp://guest:guest@some-rabbit:5672/ --net jrny_net jrny/consumer
```

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
