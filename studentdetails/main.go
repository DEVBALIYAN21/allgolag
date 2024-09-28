package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student structure
type Student struct {
    Name      string `json:"name"`
    RollNumber string `json:"roll_number"`
}

// Slice to hold student data
var students = []Student{
    {Name: "Dev Baliyan", RollNumber: "1BG22CS040"},
    {Name: "Abc", RollNumber: "124ABC"},
}

func main() { 
    r := gin.Default()
    r.GET("/students", getStudents)   //getting the details
    r.POST("/students", addStudent)   //adding the details

    // Start the server on port 8080
	fmt.Println("Running on port 8080");
    r.Run(":8080")
}

// Handler to get the list of students
func getStudents(c *gin.Context) {
    c.JSON(http.StatusOK, students)
}

// Handler to add a new student
func addStudent(c *gin.Context) {
    var newStudent Student

    // Bind the JSON body to the newStudent struct
    if err := c.ShouldBindJSON(&newStudent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Append the new student to the students slice
    students = append(students, newStudent)

    c.JSON(http.StatusCreated, newStudent)
}
/*
For uploading the details
$body = @{
    name = "MNOP"
    roll_number = "125"
}

Invoke-RestMethod -Method Post -Uri http://localhost:8080/students -ContentType "application/json" -Body ($body | ConvertTo-Json)
*/

/*
for Fetching the details
curl.exe  http://localhost:8080/students
*/