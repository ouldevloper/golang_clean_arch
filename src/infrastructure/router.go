package infrastructure

import (
	"go_clean/src/interfaces/api"
	"log"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	controller := api.NewUserController(NewSqlHandler())
	r.GET("/", func(cntx *gin.Context) {
		users := controller.GetUser()
		cntx.JSON(200, gin.H{
			"message": "Success",
			"data":    users,
		})
	})

	r.GET("/users", func(cntx *gin.Context) {
		users := controller.GetUser()
		cntx.JSON(200, gin.H{
			"message": "Success",
			"data":    users,
		})
	})

	r.GET("/user/:id", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.POST("/create", func(cntx *gin.Context) {
		controller.Create(cntx)
		cntx.JSON(200, gin.H{
			"message": "Success",
			"data":    nil,
		})
	})

	r.DELETE("/delete", func(cntx *gin.Context) {

		id := cntx.Query("id")
		controller.Delete(id)
		cntx.JSON(200, gin.H{
			"message": "Success",
			"data":    nil,
		})
	})

	log.Fatal(r.Run(":3001"))
}
