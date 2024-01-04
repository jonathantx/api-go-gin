package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathantx/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.Saudation)
	r.POST("/alunos", controllers.CreatNewAluno)
	r.GET("/alunos/:id", controllers.GetAlunoId)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditAluno)
	r.Run()

}
