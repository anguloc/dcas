package shorturl

import (
	"dcas/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
)


type TestSs struct {
	ID int
	Stock int
	Weight int
	UpdateTime int
	AddTime int
}

func Index(c *gin.Context) {
	log.Info("index:")
	c.HTML(http.StatusOK, "index.html", nil)
}
