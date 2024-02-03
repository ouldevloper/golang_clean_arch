package infrastructure

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.GET("/", func(cntx *gin.Context) {
		cntx.JSON(200, gin.H{
			"message": "desfsd",
		})
	})

	r.GET("/find", func(cntx *gin.Context) {
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
}
