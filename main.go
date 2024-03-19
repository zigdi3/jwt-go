package main

import (
	"jwt-dev/launcher"

	"github.com/gin-gonic/gin"
)

func init() {
	launcher.LoadEnvVariables();
	launcher.DbContext();
	launcher.SyncDataBases();
}

func main() {
 r:= gin.Default()
r.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
})
	r.Run();
}