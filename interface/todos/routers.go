package todos

import (
	"github.com/gin-gonic/gin"
	todosUseCase "layered-architecture-sample/usecase/todos"
)

func Initialize(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/", All)
		todos.GET("/:todoId", Get)
		todos.POST("/", Post)
		todos.PUT("/:todoId", Put)
		todos.DELETE("/:todoId", Delete)
	}

	todosUseCase.
		_ = todosUseCase.NewGetTodoUseCase()
}
