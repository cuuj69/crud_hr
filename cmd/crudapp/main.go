package main

import (
	"fmt"

	"github.com/cuuj69/crud_hr/internal/employee"
	"github.com/cuuj69/crud_hr/web/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)


var db *gorm.DB

func main() {
	fmt.Println("HR Employee Management System")


	initDB()

	router := gin.Default()

	//assets
	router.Static("/static", "./static")

	// html
	router.LoadHTMLGlob("/home/william/go/src/github.com/cuuj69/crud_hr/templates/*")


	// root
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	routes.InitORMRoutes(router, db)
	router.Run(":8080")
}


func initDB(){

	var err error

	dsn := "mensah:@forum500A@tcp(127.0.0.1:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
	   panic("Failed to connect to the database")

	}

	//Auto Migrate

	db.AutoMigrate(&employee.Employee{})


}



