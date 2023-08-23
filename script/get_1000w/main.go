package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	g := new(errgroup.Group)
	g.SetLimit(100)

	for i := 0; i < 1000; i++ {
		x := i
		g.Go(func() error {
			// POST JSON string
			client := resty.New()
			for k := x * 10000; k < (x+1)*10000; k++ {
				key := fmt.Sprintf("key%v", k)
				resp, err := client.R().
					Get("http://10.128.249.36:8080/api/get/" + key)
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
