package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vearne/cache-test/model"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	//bt, _ := json.Marshal(model.NewCar(11))
	//fmt.Println(string(bt))
	g := new(errgroup.Group)
	g.SetLimit(50)

	for i := 0; i < 1000; i++ {
		x := i
		g.Go(func() error {
			// POST JSON string
			client := resty.New()
			for k := x * 10000; k < (x+1)*10000; k++ {
				// No need to set content type, if you have client level setting
				key := fmt.Sprintf("key%v", k)
				fmt.Println("key", key, k)
				_, err := client.R().
					SetHeader("Content-Type", "application/json").
					SetBody(model.NewCar(k)).
					Put("http://10.128.249.36:8080/api/put/" + key)

				if err != nil {
					log.Println("err", err)
				} else {
					//fmt.Println("resp", resp.String())
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
