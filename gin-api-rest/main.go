package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":"1",
		"nome": "Gui Lima"
	})
}

func main() {
	r := gin.Default(":5000")
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run()
}