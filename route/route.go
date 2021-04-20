package route

import (
	"github.com/gin-gonic/gin"
	"dcas/service/shorturl"
)


// initRoute ...a
func InitRoute(g *gin.Engine) (error){
	register(g)

	return nil
}

func register(g *gin.Engine)  {
	g.GET("/", shorturl.Index);
	g.POST("/", shorturl.Gen);
}
