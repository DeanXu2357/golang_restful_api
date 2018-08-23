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
    FirstName string `json:"firstname"`
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
    r.GET("/person/", getAllPerson)
    r.GET("/person/:id", getPerson)
    r.POST("/person", createPerson)
    r.Run()
}

func createPerson(c *gin.Context) {
    var person Person
    c.BindJSON(&person)

    // db.Create(&person)
    // c.JSON(200, person)

    if dbm := db.Create(&person); dbm.Error != nil {
        c.AbortWithStatus(404)
        fmt.Println(dbm.Error)
    } else {
        c.JSON(200, person)
    }
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

func getPerson(c *gin.Context) {
    id := c.Params.ByName("id")
    var person Person
    if err := db.Where("id = ?", id).First(&person).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, person)
    }
}

