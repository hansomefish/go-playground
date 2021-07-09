package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPi(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("%s error!", api)
		return
	}
	ch <- fmt.Sprintf("%s success!", api)
}

func main() {
	ch := make(chan string)
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}
	for _, api := range apis {
		go checkAPi(api, ch)
	}
	for i:=0;i<len(apis);i++{
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Println("cost", elapsed.Seconds())
}
