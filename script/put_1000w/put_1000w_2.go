package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vearne/cache-test/model"
	"log"
)

func main() {
	for k := 0; k < 1000; k++ {
		client := resty.New()
		// No need to set content type, if you have client level setting
		key := fmt.Sprintf("key%v", k)
		car := model.NewCar(k)
		bt, _ := json.Marshal(car)
		fmt.Println("key", key, string(bt))
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(car).
			Put("http://localhost:8080/api/put/" + key)

		if err != nil {
			log.Println("err", err)
		} else {
			//fmt.Println("resp", resp.String())
		}
		log.Println("k", k)
	}
}
