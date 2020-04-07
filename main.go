package main

import (
	"fmt"
	"strconv"

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

	router.GET("/show/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, _ := strconv.Atoi(n)

		user := dbSelectShow(id)
		ctx.HTML(200, "show.html", gin.H{"user": user})
	})

	router.POST("/new", func(ctx *gin.Context) {
		name := ctx.PostForm("text")
		dbInsert(name)
		ctx.Redirect(302, "/")
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, _ := strconv.Atoi(n)
		fmt.Println(id)
		name := ctx.PostForm("text")
		dbUpdate(id, name)
		ctx.Redirect(302, "/")
	})

	router.Run()
}

//  SELECT処理（全件）
func dbSelect(users *[]User) {
	db := gormConnect()
	defer db.Close()

	db.Find(&users)
}

//  SELECT処理（1件）
func dbSelectShow(id int) User {
	db := gormConnect()
	defer db.Close()

	var user User
	db.First(&user, id)
	return user
}

// ISNERT処理
func dbInsert(name string) {
	db := gormConnect()
	defer db.Close()

	db.Create(&User{Name: name})
}

// Update処理
func dbUpdate(id int, name string) {
	db := gormConnect()
	defer db.Close()

	var user User
	db.First(&user, id)
	user.Name = name
	db.Save(&user)
}
