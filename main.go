package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vearne/cache-test/model"
	"github.com/vearne/coolcache"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":9090", nil))
	}()
	cache := coolcache.NewCache(
		coolcache.WithName("mycache"),
		coolcache.WithCapacity(20000000),
		coolcache.WithShardNumber(1000),
	)
	r := gin.Default()
	g := r.Group("/api")
	g.GET("/get/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		value := cache.Get(keyStr)
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})
	g.PUT("/put/:key", func(c *gin.Context) {
		keyStr := c.Param("key")
		var car model.Car
		err := c.BindJSON(&car)
		if err != nil {
			fmt.Println("error", err)
		} else {
			cache.Set(keyStr, &car, 30*time.Minute)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":8080")
}
