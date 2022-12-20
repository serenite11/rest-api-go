package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Song struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}
type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var songs = []Song{
	{Id: "1",
		Name: "Fire trade",
		Author: &Author{
			Id:   "1",
			Name: "Drake",
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/songs", getAllSongs)
	router.GET("/songs/:id", getSongById)
	router.POST("/songs/create", createSong)
	router.DELETE("songs/delete/:id", deleteSong)
	router.PUT("songs/update/:id", updateSong)
	router.Run(":8080")
}

func getAllSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}
func createSong(c *gin.Context) {
	var newsong Song
	if err := c.BindJSON(&newsong); err != nil {
		return
	}
	songs = append(songs, newsong)
	c.IndentedJSON(http.StatusCreated, newsong)
}

func getSongById(c *gin.Context) {
	id := c.Param("id")
	for _, value := range songs {
		if value.Id == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this song not found in songs"})
}

func deleteSong(c *gin.Context) {
	id := c.Param("id")
	for index, value := range songs {
		if value.Id == id {
			songs = append(songs[:index], songs[index+1:]...)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this song not found in songs"})
}
func updateSong(c *gin.Context) {
	var newsong Song
	if err := c.BindJSON(&newsong); err != nil {
		return
	}
	for index, value := range songs {
		if value.Id == newsong.Id {
			songs[index] = newsong
			c.IndentedJSON(http.StatusOK, newsong)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this song not found in songs"})
}
