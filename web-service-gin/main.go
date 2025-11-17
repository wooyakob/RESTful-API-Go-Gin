package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Declaration of an album struct
// Used to store album data in memory

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// assign handler function to endpoint path
// GET function to associate GET HTTP method and /albums path with handler function
// attach router to http.Server and start it
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	

	router.Run("localhost:8080")
}

// Implement Endpoint
// Logic to prepare response
// Code to map request path to logic
// Reverse of runtime execution, dependencies first, then code that depends on them

// getAlbums function creates JSON from slice of album structs, writing JSON into response
// gin Context param, carries request details, validates and serializes JSON
// calls context indented JSON to serialize struct into JSON and add to response
// function first agument is HTTP status code to send to the client
// indented form of JSON easier to work with when debugging and size diff is small from using context.JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Write handler to add a new item
// Client makes POST request at /albums - add album described in request body to existing albums data
// Logic to add new album to existing list
// Code to route POST request to logic

// postAlbums adds album from JSON received in request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call bindjson to bind received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Handler to return specific item
// Client makes GET request, return album whose ID matches ID path parameter

// getAlbumByID locates album whose ID value matches the id
// param sent by client, returns album as response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over albums list looking for id value that matches
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry, the album you're searching for wasn't found"})
}
