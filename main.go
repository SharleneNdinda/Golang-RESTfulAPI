package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// represents data about a record album
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// data we'll us too start with
var albums = []album{
    {ID: "1", Title: "Sweetner", Artist: "Ariana Grande", Price: 56.99},
    {ID: "2", Title: "Problem", Artist: "Ariana Grande", Price: 17.99},
    {ID: "3", Title: "My Everything", Artist: "Ariana Grande", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("v1/albums", getAlbums)
	router.GET("v1/albums/:id", getAlbumsbyID)
	router.POST("v1/albums", postAlbums)
	router.Run("localhost:8080")
}

// list albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// add item to albums 
func postAlbums(c *gin.Context) {
	var newAlbum album 

	// calls BindJson to bind incoming Json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to the data slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// return a specific item
func getAlbumsbyID(c *gin.Context) {
	id := c.Param("id")
;
	// loop through list to find match
	for _, a := range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}