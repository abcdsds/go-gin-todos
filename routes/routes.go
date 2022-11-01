package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateTodo)
		v1.GET("todo/:id", controllers.GetTodo)
		v1.PUT("todo/:id", controllers.UpdateTodo)
		v1.DELETE("todo/:id", controllers.DeleteTodo)
	}
	return r
}
