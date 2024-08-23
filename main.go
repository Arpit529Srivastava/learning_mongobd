package main

import (
	"errors"
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

// Fetch all movies
func GetAllMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, film)
}

// Fetch movie by ID
func moviesByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := GetMoviesById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}

// Fetch movie by name
func moviesByName(c *gin.Context) {
	name := c.Param("name")
	movie, err := GetMoviesByName(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}

// Get movie by ID
func GetMoviesById(id string) (*movies, error) {
	for i, j := range film {
		if j.Id == id {
			return &film[i], nil
		}
	}
	return nil, errors.New("movie not found")
}

// a extra function for deleting the movie
func GetMoviesById_delete(id string) (int, error) {
	for i, j := range film {
		if j.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("movie not found")
}

// Get movie by name
func GetMoviesByName(name string) (*movies, error) {
	for i, j := range film {
		if j.Name == name {
			return &film[i], nil
		}
	}
	return nil, errors.New("movie not found")
}

// for deleting the movie by using their id
func deleteMovieByID(c *gin.Context) {
	id := c.Param("id")
	index, err := GetMoviesById_delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}

	// Remove the movie from the slice
	film = append(film[:index], film[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "movie deleted"})
}

// Define the routes for the handler
func main() {
	router := gin.Default()
	router.GET("/movies", GetAllMovies)
	router.GET("/movies/id/:id", moviesByID)
	router.GET("/movies/name/:name", moviesByName)
	router.DELETE("/movies/id/:id", deleteMovieByID)
	router.Run("localhost:9090")
}
