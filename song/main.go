package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album

type album struct {
	ID     string  `json:"id"`
	Tittle string  `json:"tittle"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Tittle: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Tittle: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Tittle: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("127.0.0.1:9090")
}

// get albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums

func postAlbums(c *gin.Context) {
	var newAlbum album

	//call BindJSON to bind the received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by client, the return that album as a response.

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the lis of albums, looking for an albums whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
