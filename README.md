# RESTful-API-Go-Gin

## Tutorial

This project follows the [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin) from the official Go documentation.

## Project

[Gin](https://github.com/gin-gonic/gin) to route requests. 

Retrieve request details.

Marshal JSON for responses.

[RESTful API](https://aws.amazon.com/what-is/restful-api/) server with two endpoints.

Repository of data about vintage jazz records.

## API

Access to store selling vintage recordings on vinyl.

Endpoints for a client to get and add albums for users.

Endpoints:

* **/albums**

  * GET - list of all albums, returned as JSON.
  * POST - add new album from request data sent as JSON.
* **/albums/:id**

  * GET - get album by ID, returning album data as JSON.

Make a request to web service from CLI

```
$ curl http://localhost:8080/albums
```

**POST Album:**

```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

**GET Albums:**
curl http://localhost:8080/albums
    --header "Content-Type: application/json"
    --request "GET"

```
GET Album By ID 2:

curl http://localhost:8080/albums/2
```


## Data Structures

For handling data.

Data will be stored in memory to keep it simple but would normally interact with a database.

Storing the data in memory means the albums will be lost each time you stop the server, then recreated when you start it.

## Dependencies

```
    "net/http"

    "github.com/gin-gonic/gin"
```

go get .

## Run

/web-service-gin

go run .
