package app

import "github.com/gin-gonic/gin"

func (app *App) list(c *gin.Context) {
	c.String(200, "List all tasks")
}

func (app *App) get(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Get task with id=%s", id)
}

func (app *App) create(c *gin.Context) {
	c.String(200, "Create new task")
}

func (app *App) edit(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Edit task ith id=%s", id)
}

func (app *App) toggleStatus(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Toggle status of the task with id=%s", id)
}

func (app *App) remove(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "Delete task with id=%s", id)
}
