package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/go-devtools/handler"
)

var post = flag.String("p", "8000", "port")

func main() {
	flag.Parse()

	server := gin.New()
	handler.Route(server)

	err := server.Run(":" + *post)
	if err != nil {
		panic(err)
	}
}
