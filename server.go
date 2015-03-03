package main

import (
	"github.com/devinshively/go-gonic-archetype/api"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
	"runtime"
	"os"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()

	documentRouter := r.Group("/api/v1")
	for _, route := range api.DocumentRoutes {
		documentRouter.Handle(route.Method, route.Path, []gin.HandlerFunc{route.Handler})
	}

	r.Use(static.Serve("/", static.LocalFile("static", false)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":"+port)
}
