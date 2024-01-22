package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/web/routes"
)

func main() {
	fmt.Println("HR Employee Management System")

	// Create a new Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router)

	// Run the server on port 8080
	router.Run(":8080")
}
