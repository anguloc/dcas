package shorturl

import (
	"dcas/internal/dao"
	"dcas/internal/log"
	"github.com/gin-gonic/gin"
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
	//c.HTML(http.StatusOK, "index.html", nil)

	err := test()
	if err != nil {
		return
	}
}

func test() error {
	//res := []string{}
	//dao.DB.Exec("select 1")
	var ids []int
	//dao.DB.Model(&TestSs{}).Limit(10).Find()
	//dao.DB.Select(res, "show tables")
	dao.DB.Pluck("id", &ids)

	//dao.DB.Exec("show tables")
	//err := dao.DB.Debug().QueryFields("show tables;").Row().Scan(&res)

	//fmt.Println(res)
	//return err
	return nil
}
