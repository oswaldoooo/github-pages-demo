package main

import (
	"backend/route"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	rinfo, err := BuildRuntimeInfo(File("./config.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "build runtime info error "+err.Error())
		return
	}
	errorlogger := log.New(rinfo.loggerout, "", log.Lshortfile|log.Ltime)
	backendlog := log.New(rinfo.backendLoggerOut, "", log.Lshortfile|log.Ltime)
	bv1 := route.NewBlogV1(rinfo.db, rinfo.cache, errorlogger)
	backend := route.NewBackend(rinfo.db, rinfo.cache, backendlog)
	eng := gin.Default()
	eng.Use(Cors())
	bv1.RegisterRouter(eng.Group("/v1"))
	backend.RegisterRouter(eng.Group("/backend", backend.ParseToken))
	log.Fatalln(http.ListenAndServe(":8080", eng))
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
