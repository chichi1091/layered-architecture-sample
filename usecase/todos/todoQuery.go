package todos

import (
	"layered-architecture-sample/domain/todos"
	infraTodos "layered-architecture-sample/infrastructure/query/todos"
)

type TodoQuery interface {
	GetTodoDetail(todoId todos.TodoId) *GetTodoDTO
}

func NewTodoQuery() *TodoQuery {
	impl := infraTodos.NewTodoQueryImpl()
	return impl
}
