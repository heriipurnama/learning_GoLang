# Rest API

Go- GIN

    1. GET("/albums")
        host:port/router
        127.0.0.1:9090/albums
    2. GET("/albums/:id")
        host:port/router/:id
        127.0.0.1:9090/albums/1

    3. POST("/albums")
        127.0.0.1:9090/albums

    body JSON
    ```
    {
    "ID": "4",
    "Tittle": "Blue Train",
    "Artist": "John Coltrane",
    "Price": 56.99
    }
    ```

# Command Run with

1. go mod vendor --> to download dependencies
2. go run main.go
