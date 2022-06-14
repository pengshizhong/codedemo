package main

import (
	"net/http"
	"io"
	"log"
	"time"
)

func main()  {
	req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com/", nil)
	cli := http.Client{
		Timeout: time.Millisecond * 100000, // Set 10ms timeout.
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Panicf("%T: %+v", err, err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	log.Printf("%s", len(b))
}