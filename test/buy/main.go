package main

import (
	"fmt"
	"sync"

	httpx "github.com/liuxiaobopro/gobox/http"
)

func main() {
	var (
		url1       = "http://localhost:8080/busyBuy"
		url2       = "http://localhost:8080/goodsInfo"
		requestNum = 10
	)

	var wg sync.WaitGroup
	wg.Add(requestNum)

	for i := 0; i < requestNum; i++ {
		go func(forNum int) {
			defer wg.Done()
			httpCli := httpx.Client{Url: url1}
			resp, err := httpCli.Get()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(resp))
		}(i)
	}

	wg.Wait()

	httpCli := httpx.Client{Url: url2}
	resp, err := httpCli.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

	fmt.Println("request over")
}
