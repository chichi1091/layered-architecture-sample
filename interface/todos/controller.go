package todos

import (
	"github.com/gin-gonic/gin"
	"layered-architecture-sample/usecase/todos"
	"net/http"
)

func All(c *gin.Context) {

}

func Get(c *gin.Context) {
	id := c.Param("todoId")

	command := todos.NewGetTodoCommand(id)
	todos.NewGetTodoUseCase().Handle(command)

	c.JSON(http.StatusOK, gin.H{
		"todoId": todoId,
	})
}

func Post(c *gin.Context) {

}

func Put(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
