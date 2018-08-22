package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        //c.String(200, "hellow world")
        c.JSON(200, gin.H{
            "message": "hello world",
        })
    })
    r.Run()
}

