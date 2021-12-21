package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	f, _ := os.Create("./gin.log")
	log.Out = f

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out

}
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Warn("A group of walrus emerges from the ocean")
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
