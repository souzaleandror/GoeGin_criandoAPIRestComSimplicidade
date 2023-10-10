package main

import (
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gui", CPF: "12312312312", RG: "123123123"},
		{Nome: "Gui2", CPF: "XXXXXXXXXX", RG: "YYYYYYYYYY"},
	}
	routes.HandleRequest()
}
