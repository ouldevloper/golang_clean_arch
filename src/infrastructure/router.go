package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.GET("/users", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.GET("/user/:id", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.POST("/create", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.DELETE("/delete", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	log.Fatal(r.Run(":3001"))
}
