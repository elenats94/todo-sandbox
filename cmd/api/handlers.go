package main

import "github.com/gin-gonic/gin"

func list(c *gin.Context) {
	c.String(200, "List all tasks")
}

func get(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Get task with id=%s", id)
}

func create(c *gin.Context) {
	c.String(200, "Create new task")
}

func edit(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Edit task ith id=%s", id)
}

func toggleStatus(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Toggle status of the task with id=%s", id)
}

func remove(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Delete task with id=%s", id)
}
