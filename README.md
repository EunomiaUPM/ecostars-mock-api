# Ecostars Mock Server

This repository contains a mocking framework for a service based on the Ecostars systems. On one hand, it contains a static API to return entities related to the productive activity of Ecostars. On the other hand, it contains a PubSub system for entities that are updated in time windows with an asynchronous notification system.

The project is developed in Go 1.24. The idea is for it to serve as a playground to gradually generate a canonical API for other projects, as well as to experiment with different transport systems.

## Usage

Here is the usage tutorial

### Static API Deployment

To deploy the static API:

```bash
go run ./cmd fake     # to mock entities
go run ./cmd serve    # to start the server
```

Afterward, it can be tested:

```bash
curl --location 'http://localhost:8081/hotels'
# [
#    {
#        "name": "distinctio",
#        "address": "Dicta qui commodi sequi beatae maiores.",
#        "city": "Los Angeles",
#        "measures": [
#            {
#                "ID": 1,
#                "CreatedAt": "2025-10-28T14:00:54.186716Z",
#                "UpdatedAt": "2025-10-28T14:00:54.186716Z",
#               "DeletedAt": null,
#                "hotel_id": 1,
#                "year": 2023,
#                "rooms": 1891,
#.               ...
```

### Dynamic API Deployment

To deploy the dynamic API:

```bash
go run ./cmd metrics

# 2025/10/28 15:06:41 Enqueued notification job for metric item ID 5
# 2025/10/28 15:06:41 updated metric item ID 5: new value 91.00
# 2025/10/28 15:06:41 Enqueued notification job for metric item ID 6
# 2025/10/28 15:06:41 updated metric item ID 6: new value 27.66
# 2025/10/28 15:06:44 Enqueued notification job for metric item ID 4
# 2025/10/28 15:06:44 updated metric item ID 4: new value 33.66
# 2025/10/28 15:06:44 Enqueued notification job for metric item ID 5
# ...
```

To subscribe, simply:

```bash
curl --location 'http://localhost:8082/subscriptions/subscribe' \
--header 'Content-Type: application/json' \
--data '{
    "url": "http://localhost:1112/",
    "event_type": "hotel_water_usage"
}'

# {
#    "data": {
#        "ID": 3,
#        "CreatedAt": "2025-10-28T15:09:31.726571+01:00",
#        "UpdatedAt": "2025-10-28T15:09:31.726571+01:00",
#        "DeletedAt": null,
#        "url": "http://localhost:1112/",
#        "event_type": "hotel_water_usage"
#    },
#    "message": "Subscribed successfully"
# }
```

And to unsubscribe:

```bash
curl --location --request POST 'http://localhost:8082/subscriptions/unsubscribe/1'
```

At the notification endpoint, messages like this will arrive:

```json
{
  "id": 5,
  "item_type": "hotel_waste_generated",
  "last_value": 32.71959530690623,
  "last_measured_at": "2025-10-28T00:00:00Z"
}
```

## For use with docker

As always... just. Remember to adjust `.env` properly since docker compose services relies on this file.

```bash
# .env
DATABASE_HOST=postgres
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASS=postgres
DATABASE_NAME=postgres
STATIC_SERVER_PORT=8081
STATIC_SERVER_HOST=static-api
DYNAMIC_SERVER_PORT=8082
DYNAMIC_SERVER_HOST=dynamic-api
```

```bash
docker compose up -d
```

## Authentication & Authorization

The project uses [Keycloak](https://www.keycloak.org/) for identity management. Both `static-api` and `dynamic-api` are protected and require a valid Bearer Token.

### Keycloak Access
- **URL**: [http://localhost:8080](http://localhost:8080)
- **Admin Console**: [http://localhost:8080/admin](http://localhost:8080/admin)
- **Admin Credentials**: `admin` / `admin`

### Default Realm Configuration
The stack imports a default realm (`ecostars`) on startup with:
- **Client ID**: `ecostars-client`
- **Client Secret**: `ecostars-secret`
- **Test User**: `testuser` / `password`

### How to Authenticate
To make requests to the API, you first need to obtain an access token:

```bash
export TOKEN=$(curl -s -X POST "http://localhost:8080/realms/ecostars/protocol/openid-connect/token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "client_id=ecostars-client" \
  -d "client_secret=ecostars-secret" \
  -d "username=testuser" \
  -d "password=password" \
  -d "grant_type=password" | jq -r '.access_token')
```

Then, include the token in your requests:

```bash
curl -H "Authorization: Bearer $TOKEN" http://localhost:8081/hotels
```

## Contribute

Please open up issue to contribute
