package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type movies struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Release_date string `json:"date"`
	Director     string `json:"director"`
	Stars        int    `json:"stars"`
}

var film = []movies{
	{
		Id:           "42",
		Name:         "The Infinite Horizon",
		Release_date: "2023-07-16",
		Director:     "Jane Doe",
		Stars:        4,
	},
	{
		Id:           "17",
		Name:         "Beyond the Mist",
		Release_date: "2021-03-22",
		Director:     "John Smith",
		Stars:        3,
	},
	{
		Id:           "29",
		Name:         "Echoes of Eternity",
		Release_date: "2020-11-09",
		Director:     "Alice Johnson",
		Stars:        5,
	},
	{
		Id:           "34",
		Name:         "The Last Beacon",
		Release_date: "2022-05-12",
		Director:     "Robert Lee",
		Stars:        4,
	},
	{
		Id:           "58",
		Name:         "Whispers in the Dark",
		Release_date: "2021-08-19",
		Director:     "Emily Carter",
		Stars:        2,
	},
	{
		Id:           "63",
		Name:         "The Silent Voyager",
		Release_date: "2019-02-14",
		Director:     "Michael Brown",
		Stars:        5,
	},
	{
		Id:           "77",
		Name:         "Lost in Time",
		Release_date: "2024-01-01",
		Director:     "Sophia Green",
		Stars:        3,
	},
	{
		Id:           "82",
		Name:         "Shadows of Tomorrow",
		Release_date: "2022-09-30",
		Director:     "David Wilson",
		Stars:        4,
	},
	{
		Id:           "96",
		Name:         "The Forgotten Path",
		Release_date: "2023-06-07",
		Director:     "Lucas Martinez",
		Stars:        3,
	},
	{
		Id:           "105",
		Name:         "Rise of the Guardians",
		Release_date: "2020-12-25",
		Director:     "Olivia Clark",
		Stars:        4,
	},
}
//will fetch all the movies
func GetAllMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, film)
}
//defining the routes for the handler
func main() {
	router:= gin.Default()
	router.GET("/movies", GetAllMovies)
	router.Run("localhost:9090")
}
