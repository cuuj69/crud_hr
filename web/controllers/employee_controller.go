package controllers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/internal/employee"
	"gorm.io/gorm"																																																									
	
)

//var employees = make(map[int]*employee.Employee)
//var currentID = 1

// employee response structure
//type EmployeeResponse struct {
//	ID    int    `json:"id"`
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}


// get all employees
// func GetEmployees(c *gin.Context) {
// 	var employeeList []employee.Employee
// 	if err := db.Find(&employeeList).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employees"})
// 		return
// 	}

// 	var responseList []EmployeeResponse
// 	for _, emp := range employeeList {
// 		response := EmployeeResponse{
// 			ID:    int(emp.ID),
// 			Name:  emp.Name,
// 			Email: emp.Email,
// 		}
// 		responseList = append(responseList, response)
// 	}
// 	c.JSON(http.StatusOK, responseList)
// }

// // get employees by id
// func GetEmployeeByID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
// 		return
// 	}

// 	var emp employee.Employee
// 	if err := db.First(&emp, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
// 		return
// 	}

// 	response := EmployeeResponse{
// 		ID:    int(emp.ID),
// 		Name:  emp.Name,
// 		Email: emp.Email,
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// create employee
// func CreateEmployee(c *gin.Context) {
// 	var newEmployee employee.Employee

// 	if err := c.ShouldBindJSON(&newEmployee); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	if err := db.Create(&newEmployee).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee"})
// 		return
// 	}

// 	response := EmployeeResponse{
// 		ID:    int(newEmployee.ID),
// 		Name:  newEmployee.Name,
// 		Email: newEmployee.Email,
// 	}
// 	c.JSON(http.StatusCreated, response)
// }

// Update by id
// func UpdateEmployee(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
// 		return
// 	}

// 	var emp employee.Employee
// 	if err := db.First(&emp, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&emp); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	if err := db.Save(&emp).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating employee"})
// 		return
// 	}

// 	response := EmployeeResponse{
// 		ID:    int(emp.ID),
// 		Name:  emp.Name,
// 		Email: emp.Email,
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// DeleteEmployee deletes an employee by ID
// func DeleteEmployee(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
// 		return
// 	}

// 	if err := db.Delete(&employee.Employee{}, id).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting employee"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
// }




// ORM employee response structure
type ORMEmployeeResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Get all employees using GORM
func GetORMEmployees(c *gin.Context,db *gorm.DB) {
	var employees []employee.Employee

	if err := db.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employees from database"})
		return
	}

	var employeeList []ORMEmployeeResponse
	for _, emp := range employees {
		employeeList = append(employeeList, ORMEmployeeResponse{
			ID:    int( emp.ID),
			Name:  emp.Name,
			Email: emp.Email,
		})
	}

	c.JSON(http.StatusOK, employeeList)
}

func WrapORMGetEmployees(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		GetORMEmployees(c,db)
	}
}

// Get employee by ID using GORM
func GetORMEmployeeByID(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var emp employee.Employee

	if err := db.First(&emp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employee from database"})
		}
		return
	}

	response := ORMEmployeeResponse{
		ID:    int(emp.ID),
		Name:  emp.Name,
		Email: emp.Email,
	}

	c.JSON(http.StatusOK, response)
}

func WrapORMGetEmployeeByID(db * gorm.DB)gin.HandlerFunc {
	return func (c * gin.Context){
		GetORMEmployeeByID(c, db)
	}
}

// Create employee using GORM
func CreateORMEmployee(c *gin.Context, db *gorm.DB) {
	var newEmployee employee.Employee

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}


	if err := db.Create(&newEmployee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee in the database"})
		return
	}

	response := ORMEmployeeResponse{
		ID:    int(newEmployee.ID),
		Name:  newEmployee.Name,
		Email: newEmployee.Email,
	}

	c.JSON(http.StatusCreated, response)
}

func WrapORMCreateEmployee(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		CreateORMEmployee(c, db)
	}
}

// Update employee by ID using GORM
func UpdateORMEmployee(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var emp employee.Employee


	if err := db.First(&emp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employee from database"})
		}
		return
	}

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating employee in the database"})
		return
	}

	response := ORMEmployeeResponse{
		ID:    int(emp.ID),
		Name:  emp.Name,
		Email: emp.Email,
	}

	c.JSON(http.StatusOK, response)
}

func WrapORMUpdateEmployee(db *gorm.DB)gin.HandlerFunc{
	return func(c * gin.Context){
		UpdateORMEmployee(c, db)
	}
}

// Delete employee by ID using GORM
func DeleteORMEmployee(c *gin.Context, db * gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// if err := db.Delete(&employee.Employee{}, id).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting employee from the database"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})

	if err :=db.Unscoped().Delete(&employee.Employee{}, id).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error deleting employee from the database"})
	}
}

func WrapORMDeleteEmployee(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		DeleteORMEmployee(c, db)
	}
}