package main

import (
	"net/http"

	lib "quick-sort/lib"

	"github.com/gin-gonic/gin"
)

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func getPlaceHandler(c *gin.Context) {
	description := c.Query("description")
	place, err := lib.GetPlace(description)
	if err != nil {
		panic("failed to get place")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"place":   place,
	})
}

func addPlaceHandler(c *gin.Context) {
	var place lib.Place
	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := lib.AddPlace(place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "add success",
	})
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/health", healthCheckHandler)
	router.GET("/get-place", getPlaceHandler)
	router.POST("/add-place", addPlaceHandler)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()

}
