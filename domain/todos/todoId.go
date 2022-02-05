package todos

import (
	"fmt"
	"github.com/google/uuid"
)

type TodoId struct {
	id string
}

func NewTodoId() *TodoId {
	todoId := new(TodoId)
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	todoId.id = u.String()

	return todoId
}

func NewRestore(id string) *TodoId {
	todoId := new(TodoId)
	todoId.id = id

	return todoId
}
