package controllers

import (
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
	"github.com/cuuj69/crud_hr/internal/employee"
)

var employees = make(map[int]*employee.Employee)
var currentID = 1

// employee response structure
type EmployeeResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// get all employees
func GetEmployees(c *gin.Context) {
	var employeeList []EmployeeResponse
	for _, emp := range employees {
		employeeList = append(employeeList, EmployeeResponse{
			ID:    emp.ID,
			Name:  emp.Name,
			Email: emp.Email,
		})
	}
	c.JSON(http.StatusOK, employeeList)
}

// get employees by id
func GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if emp, ok := employees[id]; ok {
		response := EmployeeResponse{
			ID:    emp.ID,
			Name:  emp.Name,
			Email: emp.Email,
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	}
}

// create employee
func CreateEmployee(c *gin.Context) {
	var newEmployee employee.Employee

	// Log the JSON payload received
	//fmt.Println("Request JSON:", c.Request.Body)

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		//fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Log the values after binding
	//fmt.Println("New Employee:", newEmployee)

	// Assign unique id to employee
	newEmployee.ID = currentID
	currentID++

	// add employee to the map
	employees[newEmployee.ID] = &newEmployee

	response := EmployeeResponse{
		ID:    newEmployee.ID,
		Name:  newEmployee.Name,
		Email: newEmployee.Email,
	}
	c.JSON(http.StatusCreated, response)
}

// Update by id
func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if emp, ok := employees[id]; ok {
		// Bind the JSON request to the Employee struct
		if err := c.ShouldBindJSON(emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		response := EmployeeResponse{
			ID:    emp.ID,
			Name:  emp.Name,
			Email: emp.Email,
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	}
}

// DeleteEmployee deletes an employee by ID
func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Remove the employee from the map
	delete(employees, id)

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
