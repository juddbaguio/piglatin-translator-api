# Piglatin Translator API

## Notes
- [About Piglatin](https://en.wikipedia.org/wiki/Pig_Latin)

## Prerequisites
- Go v1.17
- Docker and Docker Compose (for Database instance)

## Optional
- Makefile

## Getting Started
1.) Run database through Docker by running the following command
```bash
docker compose up
```
2.) Open another terminal and run the following `make` command to start the application:
```bash
make run
```
If the system doesn't have `Makefile` installed,
```bash
go build -o bin/piglatin ./api
bin/piglatin
```
The api server runs at ```127.0.0.1:3000```.
------------------------------------------------
## API Reference:

#### Getting All Translation Requests
```http
GET / <!--- ex. 127.0.0.1:3000?page=1 --->
```
| Query Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `page` | `int` | ***Optional***. Page number of the lists |

Sample JSON Response:
```JSON
{
    "statusCode": 200,
    "data": {
        "page": 1,
        "totalPages": 1,
        "translations": [
            {
                "input": "RIDE",
                "translation": "IDERay"
            }
        ]
    }
}
```

#### Send a translation
```http
POST /translate
```
Request *JSON* Body

| Field | Type | Description
| :---- | :---- | :----------|
| `input` | `string` |  **Required**. string input you want to translate |

Sample JSON Response:
```JSON
{
    "statusCode": 201,
    "data": {
        "input": "RIDE",
        "translation": "IDERay"
    }
}
```