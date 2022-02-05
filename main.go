package main

import "layered-architecture-sample/interface/todos"

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	todos.Initialize(r)

	r.Run()
}
