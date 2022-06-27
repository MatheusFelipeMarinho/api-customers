package routes

import (
	"api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	
	r.GET("/customer", controllers.All)
	r.GET("/:nome", controllers.Congrats)
	r.POST("/customer", controllers.Create)
	r.GET("/customer/:id", controllers.FindById)
	r.DELETE("/customer/:id", controllers.Delete)
	r.PATCH("/customer/:id", controllers.Edit)
	r.GET("/customer/cpf/:cpf", controllers.FindByCpf)

	r.Run()
}
