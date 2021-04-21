package shorturl

import (
	"dcas/internal/dao"
	"dcas/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Index(c *gin.Context) {
	log.Info("index:")
	c.HTML(http.StatusOK, "index.html", nil)
	row:=dao.DB.Debug().Raw("show tables;").Row()
}
