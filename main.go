package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
)

func main() {
	r := gin.Default()
	r.GET("/busyBuy", busyBuy)

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
