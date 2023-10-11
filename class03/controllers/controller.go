package controllers

import (
	"gin-api-rest/models"
	"gin-api-rest/database"
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

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err != c.ShouldBindJSON(&aluno); err != nil {
		e.JSON(http.StatusBadRequest, gin.H {
			"Error": err.Error()
		})
		return
	}
	database.DB.Create(&aluno)
	e.JSON(http.StatusOK, aluno)
}