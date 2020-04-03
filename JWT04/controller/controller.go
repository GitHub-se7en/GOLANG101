package controller

import (
	"GOLANG101/JWT04/models"
	"GOLANG101/JWT04/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//List all todos
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := repository.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Create a Todo
func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := repository.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = repository.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Delete a Todo
func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
