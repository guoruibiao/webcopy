package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/webcopy/controllers"
)
func Init(engine *gin.Engine) {
	imageRouter := engine.Group("/image")
	{
		imageRouter.GET("/index", controllers.ImageIndex)
	}

	textRouter := engine.Group("/text")
	{
		textRouter.GET("/index", controllers.TextIndex)
		textRouter.GET("/get", controllers.TextGet)
		//textRouter.GET("/page", controllers.TextPage)
		textRouter.POST("/update", controllers.TextUpdate)
	}
}
