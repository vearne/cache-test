package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/sync/errgroup"
	"log"
)

const host = "192.168.2.100:8080"

func main() {
	g := new(errgroup.Group)
	g.SetLimit(100)

	url := fmt.Sprintf("http://%v/api/v1/get/", host)
	for i := 0; i < 1000; i++ {
		x := i
		g.Go(func() error {
			// POST JSON string
			client := resty.New()
			for k := x * 10000; k < (x+1)*10000; k++ {
				//time.Sleep(50 * time.Millisecond)
				key := fmt.Sprintf("key%v", k)
				resp, err := client.R().
					Get(url + key)
				if err != nil {
					log.Println("err", err)
				} else {
					fmt.Println("resp", resp.String())
				}
			}
			return nil
		})
		log.Println("i", i)
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Success.")
	}
}
