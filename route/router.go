package route

import (
	"CSV_COnvert_Tool/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/css", "/public/css")
	router.LoadHTMLGlob("view/*html")
	router.GET("/", controller.IndexDisplayAction)
	router.POST("/regit", controller.Regit_Date)
	return router
}
