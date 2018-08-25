[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/UlboraContentService)](https://goreportcard.com/report/github.com/Ulbora/UlboraContentService)

Ulbora Headless CMS
==============

Ulbora Headless CMS is a CMS that can provide content for use anywhere. 


## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, and DELETE. No token required for get services.)
- clientId: clientId (example 33477)



## Add Content

```
POST:
URL: http://localhost:3008/rs/content/add

Example Request
{
   "title":"new title for insert",
   "category": "books",
   "text":"some text",
   "hits": 100,
   "metaAuthorName": "ken",
   "metaDesc": "ken",
   "metaKeyWords": "ulbora, content, microservice",
   "metaRobotKeyWords": "ulbora, content, microservice"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Content

```
PUT:
URL: http://localhost:3008/rs/content/update

Example Request
{
   "id": 11,
   "title":"new title for insert for ken with oauth again",
   "category": "Books",
   "text":"some text",
   "hits": 100,
   "metaAuthorName": "ken",
   "metaDesc": "ken",
   "metaKeyWords": "ulbora, content, microservice",
   "metaRobotKeyWords": "ulbora, content, microservice"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 11
}

```



## Update Content Hits

```
PUT:
URL: http://localhost:3008/rs/content/hits

Example Request
{
   "id": 11,  
   "hits": 101
}
  
```

```
Example Response   

{
    "success": true,
    "id": 11
}

```



## Get Content

```
GET:
URL: http://localhost:3008/rs/content/get/11/403
  
```

```
Example Response   

{
    "id": 11,
    "title": "new title for insert for ken with oauth again",
    "category": "Guns",
    "createDate": "2017-08-15T01:09:36Z",
    "modifiedDate": "2017-08-15T01:13:55Z",
    "hits": 100,
    "metaAuthorName": "ken",
    "metaDesc": "ken",
    "metaKeyWords": "ulbora, content, microservice",
    "metaRobotKeyWords": "ulbora, content, microservice",
    "text": "some text",
    "clientId": 403
}

```


## Get Content By Client

```
GET:
URL: http://localhost:3008/rs/content/list/403
  
```

```
Example Response   

[
    {
        "id": 11,
        "title": "new title for insert for ken with oauth again",
        "category": "Guns",
        "createDate": "2017-08-15T01:09:36Z",
        "modifiedDate": "2017-08-15T01:13:55Z",
        "hits": 100,
        "metaAuthorName": "ken",
        "metaDesc": "ken",
        "metaKeyWords": "ulbora, content, microservice",
        "metaRobotKeyWords": "ulbora, content, microservice",
        "text": "some text",
        "clientId": 403
    },
    {
        "id": 14,
        "title": "new title for insert",
        "category": "books",
        "createDate": "2017-08-15T01:23:45Z",
        "modifiedDate": "0001-01-01T00:00:00Z",
        "hits": 100,
        "metaAuthorName": "ken",
        "metaDesc": "ken",
        "metaKeyWords": "ulbora, content, microservice",
        "metaRobotKeyWords": "ulbora, content, microservice",
        "text": "some text",
        "clientId": 403
    }
]

```


## Delete Content

```
DELETE:
URL: http://localhost:3008/rs/content/delete/14
  
```

```
Example Response   

{
    "success": true,
    "id": 15
}

```
# Docker usage
```
docker run --network=ulbora_bridge --name content --log-opt max-size=50m --env DATABASE_HOST=someHost /
 --env DATABASE_USER_NAME=someUser --env DATABASE_USER_PASSWORD=somePw --env DATABASE_NAME=ulbora_content_service /
 --env DATABASE_POOL_SIZE=5 --env OAUTH2_VALIDATION_URI=http://oauth2:8080/rs/token/validate /
 --env --restart=always -d ulboralabs/content sh
```