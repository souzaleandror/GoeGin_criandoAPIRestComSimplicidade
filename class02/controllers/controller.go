package controllers

import (
	"gin-api-rest/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"id":   "1",
	// 	"nome": "Gui Lima",
	// })
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", Tudo beleza?",
	})
}
