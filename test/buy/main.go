package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:8080/busyBuy"
	forNum := 0
	requestNum := 10

	var wg sync.WaitGroup
	wg.Add(requestNum)

	for {
		if forNum > requestNum {
			break
		}
		go func() {
			defer wg.Done()
			// get请求
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))
		}()
		forNum++
	}

	wg.Wait()
	fmt.Println("request over")
}
