# mailing-service
microservice for storing mailing details and sending emails

## Running
requirements:
docker & docker-compose
```make run-local```


endpoints:
- POST /api/messages add new mailing details
- POST /api/messages/send send emails to customers (sending is mocked)
- DELETE /api/messages/{id} delete mailing details

service has a "cron-job" that runs every 60s and deletes all mailing details that have insert_time older than 5 minutes


## Example curls

```curl -X POST localhost:8080/api/messages -d '{"email":john.doe@example.com","title":"title","content":"simple text","mailing_id":1, "insert_time": "2020-04-24T05:42:38.725412916Z"}'```

```curl -X POST localhost:8080/api/messages -d '{"email":"john.doe@example.com","title":"title","content":"simple text","mailing_id":2, "insert_time": "2020-04-24T05:42:38.725412916Z"}'```

```curl -X POST localhost:8080/api/messages -d '{"email":"john.doe@example.com","title":"title","content":"simple text","mailing_id":3, "insert_time": "2020-04-24T05:42:38.725412916Z"}'```

```curl -X POST localhost:8080/api/messages/send -d '{"mailing_id":1}'```

```curl -X DELETE localhost:8080/api/messages/{id}```
