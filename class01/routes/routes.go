package routes
import (
	"github.com/gin-gonic/gin"
)


func HandleRequest() {
	r := gin.Default(":5000")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run()
}