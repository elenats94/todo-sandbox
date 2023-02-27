package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (app *App) list(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	tasks, _ := app.storage.ListTasks(owner)
	c.JSON(http.StatusOK, tasks)
}

func (app *App) get(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	task, err := app.storage.GetTaskByID(id, owner)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (app *App) create(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}

	title, ok := data["title"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "\"title\" field missing",
		})
		return
	}

	task, _ := app.storage.CreateTask(title, owner)
	c.JSON(http.StatusCreated, task)
}

func (app *App) edit(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	var data map[string]string
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}

	title, ok := data["title"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "\"title\" field missing",
		})
		return
	}

	task, err := app.storage.UpdateTask(id, title, owner)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (app *App) toggleStatus(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	task, err := app.storage.ToggleStatus(id, owner)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (app *App) remove(c *gin.Context) {
	rawID, _ := c.Get("owner_id")
	owner := uuid.MustParse(rawID.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	task, err := app.storage.DeleteTask(id, owner)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
