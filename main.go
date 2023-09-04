package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vearne/cache-test/middleware"
	"github.com/vearne/cache-test/model"
	"github.com/vearne/coolcache"
	"github.com/vearne/golib/prom_cache"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":18080", nil))
	}()

	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// segment lock
	cache := coolcache.NewCache(
		coolcache.WithKind("mycache"),
		coolcache.WithCapacity(20000000),
		coolcache.WithShardNumber(1000),
	)

	multiMapGroup := r.Group("/api/v1")
	multiMapGroup.Use(middleware.Metric())
	multiMapGroup.GET("/get/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		value := cache.Get(keyStr)
		log.Println("get", keyStr, value)
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})
	multiMapGroup.PUT("/put/:key", func(c *gin.Context) {
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

	// single map as cache
	fixedCache := prom_cache.NewCacheWithProm("single", 10000000)

	singleMapGroup := r.Group("/api/v2")
	singleMapGroup.Use(middleware.Metric())
	singleMapGroup.GET("/get/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		value := fixedCache.Get(keyStr)
		log.Println("get", keyStr, value)
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})
	singleMapGroup.PUT("/put/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		var car model.Car
		car.CarWheel = &model.Wheel{}
		car.CarDoor = &model.Door{}
		err := c.BindJSON(&car)
		if err != nil {
			fmt.Println("error", err)
		} else {
			fixedCache.Set(keyStr, &car, 3*time.Hour)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":8080")
}
