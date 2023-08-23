package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vearne/cache-test/middleware"
	"github.com/vearne/cache-test/model"
	"github.com/vearne/coolcache"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":18080", nil))
	}()
	cache := coolcache.NewCache(
		coolcache.WithName("mycache"),
		coolcache.WithCapacity(20000000),
		coolcache.WithShardNumber(1000),
	)
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	g := r.Group("/api")
	g.Use(middleware.Metric())
	g.GET("/get/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		value := cache.Get(keyStr)
		log.Println("get", keyStr, value)
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})
	g.PUT("/put/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		var car model.Car
		car.CarWheel = &model.Wheel{}
		car.CarDoor = &model.Door{}
		err := c.BindJSON(&car)
		if err != nil {
			fmt.Println("error", err)
		} else {
			cache.Set(keyStr, &car, 3*time.Hour)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":8080")
}
