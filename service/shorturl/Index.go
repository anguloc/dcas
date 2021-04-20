package shorturl

import (
	"dcas/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Index(c *gin.Context) {
	log.Info("index:")
	c.HTML(http.StatusOK, "index.html", nil)
}
