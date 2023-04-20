package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
)

type handle struct {
	BuyLock sync.Mutex
}

func main() {
	r := gin.Default()

	var handle handle
	r.GET("/busyBuy", handle.busyBuy)
	r.GET("/goodsInfo", handle.goodsInfo)

	fmt.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

var (
	Book = GoodsInfoType{
		Id:   1,
		Name: "Book",
		Num:  100,
	}
	NowBuyNum = 0
	killNum   = 10
)

func (handle *handle) busyBuy(c *gin.Context) {
	uid := c.Query("uid")

	handle.BuyLock.Lock()
	defer handle.BuyLock.Unlock()

	if Book.Num <= 0 {
		c.JSON(http.StatusOK, respx.T{
			Code: respx.FailErrCode,
			Msg:  "商品已售罄",
		})
		return
	}

	if NowBuyNum >= killNum {
		c.JSON(http.StatusOK, respx.T{
			Code: respx.FailErrCode,
			Msg:  "秒杀已结束",
		})
		return
	}

	_ = Book.Buy()

	NowBuyNum++

	fmt.Println("uid:", uid, "buy success")
	c.JSON(http.StatusOK, respx.Succ(map[string]interface{}{
		"uid":  uid,
		"book": Book,
	}))
}

func (handle *handle) goodsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, respx.Succ(Book))
}
