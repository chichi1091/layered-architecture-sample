package todos

import "layered-architecture-sample/domain/todos"

type GetTodoUseCase interface {
	Handle(command *GetTodoCommand) *GetTodoDTO
}

type getTodoUseCase struct {
	todoQuery *TodoQuery
}

func NewGetTodoUseCase(todoQuery *TodoQuery) GetTodoUseCase {
	return getTodoUseCase{
		todoQuery: todoQuery,
	}
}

func (g getTodoUseCase) Handle(command *GetTodoCommand) *GetTodoDTO {

}

type GetTodoCommand struct {
	todoId todos.TodoId
}
func NewGetTodoCommand(id string) *GetTodoCommand {
	command := new(GetTodoCommand)
	command.todoId := todos.NewRestore(id)

	return command
}

type GetTodoDTO struct {
	todoId todos.TodoId
	title string
	content string
}