package app

import "github.com/gin-gonic/gin"

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(addr string) error {
	return app.routes().Run(addr)
}

func (app *App) routes() *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
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
