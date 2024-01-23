package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/web/routes"
)

func main() {
	fmt.Println("HR Employee Management System")

	router := gin.Default()

	// Serve static assets (styles.css, app.js, etc.)
	router.Static("/static", "./static")

	// Load HTML templates from the templates folder
	router.LoadHTMLGlob("/home/william/go/src/github.com/cuuj69/crud_hr/templates/*")


	// Define the homepage route
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	routes.InitRoutes(router)

	router.Run(":8080")
}
