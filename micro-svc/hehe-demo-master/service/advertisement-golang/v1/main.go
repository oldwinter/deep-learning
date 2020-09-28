package main

import (
	"time"
	"log"
	"math/rand"
	"github.com/gin-gonic/gin"
)

var incomingHeaders = []string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
}

// TraceMiddleware add trace header
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, h := range incomingHeaders {
			if ih, ok := c.Request.Header[h]; ok {
				log.Print(h, ih[0])
			}
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(TraceMiddleware())
	r.GET("/ad", func(c *gin.Context) {
		var fruits =  [...]string{0: "airCleanImg",1: "tshirtImg",2: "mountainClimbingImg"}
		rand.Seed(time.Now().Unix())
		//[0,3)随机值
		var index = rand.Intn(3)
		c.JSON(200, gin.H{
			"adImgName": fruits[index],
		})
	})
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Run(":3003")
}
