package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
)

var (
	GoodsInfo = GoodsInfoType{Id: 1, Name: "goods1", Num: 10}
)

func main() {
	r := gin.Default()
	r.GET("/busyBuy", busyBuy)
	r.GET("/goodsInfo", goodsInfo)

	fmt.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

type busyBuyRes struct {
	Id     int  `json:"id"`
	IsSucc bool `json:"isSucc"`
}

func busyBuy(c *gin.Context) {
	c.JSON(http.StatusOK, respx.Succ(busyBuyRes{
		Id:     1,
		IsSucc: true,
	}))
}

func goodsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, respx.Succ(GoodsInfo))
}
