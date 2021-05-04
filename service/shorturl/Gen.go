package shorturl

import (
	"dcas/config"
	"dcas/internal/dao"
	"dcas/internal/log"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

type ShortUrl struct {
	ID        int
	ShortKey  string
	Url       string
	AddTime   int32
	IsDeleted int
}

func Gen(c *gin.Context) {
	longUrl := c.PostForm("url")
	if reflect.TypeOf(longUrl).Name() != "string" || longUrl == "" {
		c.SecureJSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "url为空",
		})
		return
	}

	key,err := getKey()
	if err != nil {
		log.Error("generate key fail:" + err.Error())
		c.SecureJSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "url生成失败",
		})
		return
	}
	log.Info("Gen:req:" + longUrl + ":key:" + key)

	data := ShortUrl{ShortKey: key,Url: longUrl,AddTime: int32(time.Now().Unix())}
	res := dao.DB.Table("short_url").Create(&data)
	if res.Error != nil {
		log.Error("generate key fail:" + res.Error.Error())
		c.SecureJSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "url生成失败",
		})
		return
	}

	surl := config.Conf.ShortDomain + key;

	resp := gin.H{
		"code": 0,
		"data": gin.H{
			"url": surl,
		},
	}
	c.SecureJSON(http.StatusOK, resp)
}

func getKey()(string,error)  {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "",err
	}
	id := node.Generate().Int64()
	return idTo62(id),nil
}

func idTo62(num int64) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	res := ""
	for num > 0 {
		r := num % 62
		num /= 62
		res = string(str[r]) + res
	}
	return res
}
