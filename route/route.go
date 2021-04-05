package route

import (
	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
)

type Server struct {
	GinEngine *gin.Engine
}

// initRoute ...a
func Init() (*Server,error){
	server := new(Server)

	server.GinEngine = gin.Default()
	server.GinEngine.Use(gin.Recovery())
	//server.GinEngine.Use(gin.Logger())

	return server,nil
}
