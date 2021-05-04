package main

import (
	"dcas/config"
	"dcas/internal/dao"
	"dcas/internal/log"
	"dcas/route"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	dao.InitDB()

	gin.SetMode(gin.DebugMode)

	g := gin.Default()
	g.Use(gin.Recovery())

	g.Static("static", "public/static")
	g.LoadHTMLFiles("public/index.html")

	err = route.InitRoute(g)
	if err != nil {
		log.Fatal("init route fail:" + err.Error())
		return
	}


	port := config.Conf.Server.Port

	if port == "" {
		log.Fatal("invalid port")
		return
	}

	err = g.Run("localhost:" + port)
	if err != nil {
		log.Fatal("init route fail:" + err.Error())
		return
	}
}
