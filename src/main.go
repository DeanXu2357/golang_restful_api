package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct {
    ID uint `json:"id"`
    FisrtName string `json:"firstname"`
    LastName string `json:"lastname"`
}

func main() {

    db, err = gorm.Open("sqlite3", "./gorm.db")

    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    db.AutoMigrate(&Person{})

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        //c.String(200, "hellow world")
        c.JSON(200, gin.H{
            "message": "hello world",
        })
    })
    r.GET("/person", getAllPerson)
    r.Run()
}

func getAllPerson(c *gin.Context) {
    var people []Person
    if err := db.Find(&people).Error; err == nil {
        c.JSON(200, people)
    } else {
        c.AbortWithStatus(404)
        fmt.Println(err)
    }
}

