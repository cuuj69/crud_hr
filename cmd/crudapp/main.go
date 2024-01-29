package main

import (
	"fmt"

	"github.com/cuuj69/crud_hr/internal/employee"
	"github.com/cuuj69/crud_hr/web/routes"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"os"
)


var db *gorm.DB

func main() {
	fmt.Println("HR Employee Management System")

	//load .env

	if err := godotenv.Load(); err != nil{
		panic("Error loading .env file")
	}

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

	// dsn := "mensah:@forum500A@tcp(127.0.0.1:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
	   panic("Failed to connect to the database")

	}

	//Auto Migrate

	db.AutoMigrate(&employee.Employee{})


}



