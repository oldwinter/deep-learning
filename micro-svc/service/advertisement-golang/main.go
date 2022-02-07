package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
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

var (
	address     = "0.0.0.0:3007"
)

func main() {
	count := 0

	//从环境变量获取配置信息
	newsServerName,isExist := os.LookupEnv("NEWS_SERVER_NAME")
	if isExist {
		address = newsServerName
	}else {
		address = "0.0.0.0:3007"
	}


	r := gin.Default()
	r.Use(TraceMiddleware())
	r.GET("/ad", func(c *gin.Context) {
		count++

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		cgrpc := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		name := strconv.Itoa(count)
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := cgrpc.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)

		var fruits = [...]string{0: "airCleanImg", 1: "tshirtImg", 2: "mountainClimbingImg"}
		rand.Seed(time.Now().Unix())
		//[0,3)随机值
		var index = rand.Intn(3)
		c.JSON(200, gin.H{
			"adImgName":    fruits[index],
			"adSlogan":     r.Message,
			"requestCount": count,
		})

	})
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Run(":3003")
}
