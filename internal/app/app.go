package app

import (
	"github.com/gin-gonic/gin"
	"todo-sandbox/internal/middleware"
	"todo-sandbox/internal/store"
)

type App struct {
	storage *store.Store
}

func NewApp() *App {
	return &App{
		storage: store.NewStore(),
	}
}

func (app *App) Run(addr string) error {
	return app.routes().Run(addr)
}

func (app *App) routes() *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	tasks.Use(middleware.Identification())
	{
		tasks.GET("", app.list)
		tasks.POST("", app.create)
		single := tasks.Group("/:id")
		{
			single.GET("", app.get)
			single.PUT("", app.edit)
			single.PATCH("", app.toggleStatus)
			single.DELETE("", app.remove)
		}
	}

	return r
}
