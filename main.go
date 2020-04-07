package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int
	Name string
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

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(ctx *gin.Context) {
		users := []User{}
		dbSelect(&users)
		ctx.HTML(200, "index.html", gin.H{"users": users})
	})

	router.POST("/new", func(ctx *gin.Context) {
		name := ctx.PostForm("text")
		dbInsert(name)
		ctx.Redirect(302, "/")
	})

	router.Run()
}

//  SELECT処理
func dbSelect(users *[]User) {
	db := gormConnect()
	defer db.Close()

	db.Find(&users)
}

// ISNERT処理
func dbInsert(name string) {
	db := gormConnect()
	defer db.Close()

	db.Create(&User{Name: name})
}
