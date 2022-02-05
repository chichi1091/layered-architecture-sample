package todos

import (
	domainTodo "layered-architecture-sample/domain/todos"
	useCase "layered-architecture-sample/usecase/todos"
)

type TodoQueryImpl struct {
}

func NewTodoQueryImpl() *TodoQueryImpl {
	var query TodoQueryImpl
	return &query
}

func (q useCase.TodoQuery) GetTodoDetail(todoId domainTodo.TodoId) *useCase.GetTodoDTO {
	if q == nil {

	}

	return useCase.GetTodoDTO{
		domainTodo.NewTodoId(),
		"title",
		"content",
	}
}
