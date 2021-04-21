package main

import (
	"dcas/config"
	"dcas/route"
	"dcas/utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	g := gin.Default()
	g.Use(gin.Recovery())

	g.LoadHTMLGlob("public/*")

	err := route.InitRoute(g)
	if err != nil {
		log.Fatal("init route fail:" + err.Error())
		return
	}

	port := config.Conf.Server.Port

	if port == "" {
		log.Fatal("invalid port")
		return
	}

	err = g.Run(":" + port)
	if err != nil {
		log.Fatal("init route fail:" + err.Error())
		return
	}
}
