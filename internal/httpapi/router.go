
package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}


func Router() http.Handler {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/health", health)
	return router

}
