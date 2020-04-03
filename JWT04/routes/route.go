package routes

import (
	"GOLANG101/JWT04/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controller.GetTodos)
		v1.POST("todo", controller.CreateATodo)
		v1.GET("todo/:id", controller.GetATodo)
		v1.PUT("todo/:id", controller.UpdateATodo)
		v1.DELETE("todo/:id", controller.DeleteATodo)
	}
	return r
}
