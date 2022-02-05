package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	todoId := c.Param("todoId")
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
