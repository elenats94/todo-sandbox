package app

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"}}))

	tasks := r.Group("/tasks")
	tasks.Use(middleware.CookieController())
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
