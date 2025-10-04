// 代码生成时间: 2025-10-05 03:36:25
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/labstack/echo"
)

// Student represents a student's profile
type Student struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Email    string `json:"email"`
    Grade    string `json:"grade"`
}

// studentStore simulates a database for storing student profiles
var studentStore = make(map[string]Student)

// indexHandler returns a list of all student profiles
func indexHandler(c echo.Context) error {
    var profiles []Student
    for _, profile := range studentStore {
        profiles = append(profiles, profile)
    }
    return c.JSON(http.StatusOK, profiles)
}

// getHandler retrieves a student profile by ID
func getHandler(c echo.Context) error {
    profileID := c.Param("profile_id")
    profile, exists := studentStore[profileID]
    if !exists {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": fmt.Sprintf("Student with ID %s not found", profileID),
        })
    }
    return c.JSON(http.StatusOK, profile)
}

// postHandler creates a new student profile
func postHandler(c echo.Context) error {
    var newProfile Student
    if err := c.Bind(&newProfile); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }
    studentStore[newProfile.ID] = newProfile
    return c.JSON(http.StatusCreated, newProfile)
}

// putHandler updates an existing student profile
func putHandler(c echo.Context) error {
    profileID := c.Param("profile_id")
    var updatedProfile Student
    if err := c.Bind(&updatedProfile); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }
    if _, exists := studentStore[profileID]; !exists {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": fmt.Sprintf("Student with ID %s not found", profileID),
        })
    }
    studentStore[profileID] = updatedProfile
    return c.JSON(http.StatusOK, updatedProfile)
}

// deleteHandler deletes a student profile by ID
func deleteHandler(c echo.Context) error {
    profileID := c.Param("profile_id")
    if _, exists := studentStore[profileID]; !exists {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": fmt.Sprintf("Student with ID %s not found", profileID),
        })
    }
    delete(studentStore, profileID)
    return c.NoContent(http.StatusNoContent)
}

func main() {
    e := echo.New()
    e.GET("/students", indexHandler)
    e.GET("/students/:profile_id", getHandler)
    e.POST("/students", postHandler)
    e.PUT("/students/:profile_id", putHandler)
    e.DELETE("/students/:profile_id", deleteHandler)
    
    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}