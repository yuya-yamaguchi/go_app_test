package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type user struct {
	Id   int    `json:id`
	Name string `json:name`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	// DB接続
	db := gormConnect()
	// DBクローズ
	defer db.Close()

	users := []user{}
	// db.First(&user, "name=?", "taro")
	db.Find(&users)

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"users": users})
	})

	router.Run()
}
