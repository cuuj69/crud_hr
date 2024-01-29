//usr/bin/env go run "$0" "$@"; exit "$?"

package routes

import (
	"github.com/cuuj69/crud_hr/web/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func InitRoutes(router *gin.Engine) {
// 	// Employee routes
// 	employeeGroup := router.Group("/employees")
// 	{
// 		employeeGroup.GET("/", controllers.GetEmployees)
// 		employeeGroup.GET("/:id", controllers.GetEmployeeByID)
// 		employeeGroup.POST("/", controllers.CreateEmployee)
// 		employeeGroup.PUT("/:id", controllers.UpdateEmployee)
// 		employeeGroup.DELETE("/:id", controllers.DeleteEmployee)
// 	}

// 	initORMRoutes(router)
// }




func InitORMRoutes(router *gin.Engine, db* gorm.DB) {
	// Employee routes using ORM (GORM)
	employeeGroup := router.Group("/orm/employees")
	{
		employeeGroup.GET("/", controllers.WrapORMGetEmployees(db))
		employeeGroup.GET("/:id", controllers.WrapORMGetEmployeeByID(db))
		employeeGroup.POST("/", controllers.WrapORMCreateEmployee(db))
		employeeGroup.PUT("/:id", controllers.WrapORMUpdateEmployee(db))
		employeeGroup.DELETE("/:id", controllers.WrapORMDeleteEmployee(db))
	}
}

