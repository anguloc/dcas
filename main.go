package main

import (
	"github.com/gin-gonic/gin"

	"dcas/service/shorturl"
)

func main() {
	gin.SetMode(gin.DebugMode)


	shorturl.Index()

	//return
	//server,err := route.Init()
	//if err != nil {
	//	return err
	//}
	//server.GinEngine.Run()
	//
	//println(server)
	//
	//return
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
}