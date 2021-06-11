package initialize

import (
	"github.com/gin-gonic/gin"
)

/*NewEngine used for create new engine*/
func NewEngine(config string) *gin.Engine {
	setConfig(config)
	router := gin.Default()
	setBlueprintApp(router)
	setErrorHandler(router)
	return router
}
