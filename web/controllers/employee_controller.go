//usr/bin/env go run "$0" "$@"; exit "$?"

package controllers

import (
	
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/cuuj69/crud_hr/internal/employee"


)

  
var employees = make(map[int]*employee.Employee)
var currentID = 1

// get all employees
func GetEmployees(c *gin.Context) {
	
	c.JSON(http.StatusOK, employees)
}

// get employee by id
func GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	
	if emp, ok := employees[id]; ok {
		c.JSON(http.StatusOK, emp)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	}
}

// create employee
func CreateEmployee(c *gin.Context) {
	var newEmployee employee.Employee

	
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Assign unique id to employee
	newEmployee.ID = currentID
	currentID++

	// add employee
	employees[newEmployee.ID] = &newEmployee

	c.JSON(http.StatusCreated, newEmployee)
}

// Update by id

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	//
	if emp, ok := employees[id]; ok {
		// Bind the JSON request to the Employee struct
		if err := c.ShouldBindJSON(emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		c.JSON(http.StatusOK, emp)
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

