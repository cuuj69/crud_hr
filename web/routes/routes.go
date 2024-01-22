//usr/bin/env go run "$0" "$@"; exit "$?"

package routes


import (
	
	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/web/controllers"

)


func InitRoutes(router *gin.Engine) {
	// Employee routes
	employeeGroup := router.Group("/employees")
	{
		employeeGroup.GET("/", controllers.GetEmployees)
		employeeGroup.GET("/:id", controllers.GetEmployeeByID)
		employeeGroup.POST("/", controllers.CreateEmployee)
		employeeGroup.PUT("/:id", controllers.UpdateEmployee)
		employeeGroup.DELETE("/:id", controllers.DeleteEmployee)
	}
}
