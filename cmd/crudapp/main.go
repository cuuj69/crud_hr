//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/web/routes"
)

func main() {
	fmt.Println("HR Employee Management System")

	router := gin.Default()
	
	routes.InitRoutes(router)

	router.Run(":8080")
}
