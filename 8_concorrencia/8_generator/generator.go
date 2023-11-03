package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			time.Sleep(time.Second * 5)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return ch
}

func main() {
	t1 := titulo("https://www.google.com", "https://github.com")
	t2 := titulo("https://www.uol.com.br", "https://www.instagram.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
