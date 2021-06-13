package initialize

import (
	"frankie/universal/sdk/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setBlueprintApp(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// router.POST("/isgood", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	ctrl := controllers.NewValidation()
	router.POST("/isgood", ctrl.Validate)

}

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
