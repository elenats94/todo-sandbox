package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.GET("", list)
		tasks.POST("", create)
		single := tasks.Group("/:id")
		{
			single.GET("", get)
			single.PUT("", edit)
			single.PATCH("", toggleStatus)
			single.DELETE("", remove)
		}
	}

	log.Fatal(router.Run())
}
