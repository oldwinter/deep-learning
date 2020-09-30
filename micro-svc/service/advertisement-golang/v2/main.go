package main

import (
	"time"
	"log"
	"context"
	"os"
	"math/rand"
	"strconv"
	"github.com/gin-gonic/gin"

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

const (
	//本地调试
	address     = "0.0.0.0:3007"
	defaultName = "world"
)

func main() {
	count := 0

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
		
		
		var fruits =  [...]string{0: "airCleanImg",1: "tshirtImg",2: "mountainClimbingImg"}
		rand.Seed(time.Now().Unix())
		//[0,3)随机值
		var index = rand.Intn(3)
		c.JSON(200, gin.H{
			"adImgName": fruits[index],
			"adSlogan": r.Message,
			"requestCount": count,
		})
		


	})
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Run(":3003")
}
