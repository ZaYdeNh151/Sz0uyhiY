// 代码生成时间: 2025-10-08 02:37:21
package main

import (
    "encoding/json"
    "net/http"
    "strings"
    "log"

    "github.com/labstack/echo"
)

// SensorData represents the structure of sensor data.
type SensorData struct {
    Timestamp string `json:"timestamp"`
    Value     float64 `json:"value"`
}

func main() {
    e := echo.New()
    
    // Define a route for collecting sensor data.
    e.POST("/collect", collectSensorData)
    
    // Start the Echo server.
    log.Fatal(e.Start(":8080"))
}

// collectSensorData is an Echo handler function for collecting sensor data.
func collectSensorData(c echo.Context) error {
    // Create a new instance of SensorData to hold the incoming data.
    sd := SensorData{}
    
    // Bind the incoming JSON to the SensorData struct.
    if err := c.Bind(&sd); err != nil {
        return err
    }
    
    // Check if the received data is valid.
    if sd.Timestamp == "" || sd.Value == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid sensor data")
    }
    
    // Here you would typically add logic to save the sensor data to a database or perform further processing.
    // For demonstration, we'll just log the received data.
    log.Printf("Received sensor data: %+v", sd)
    
    // Respond with a success message and the received data.
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Sensor data collected successfully",
        "data": sd.Timestamp,
    })
}
