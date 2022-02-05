package todos

import "github.com/gin-gonic/gin"

func Initialize(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/:todoId", Get)
		todos.POST("/", Post)
		todos.PUT("/:todoId", Put)
		todos.DELETE("/:todoId", Delete)
	}
}
