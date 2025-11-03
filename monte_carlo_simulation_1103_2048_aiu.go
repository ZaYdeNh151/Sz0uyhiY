// 代码生成时间: 2025-11-03 20:48:54
package main

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "net/http"
    "time"

    "github.com/labstack/echo/v4"
)

// MonteCarloSimulation represents the parameters for the Monte Carlo simulation.
type MonteCarloSimulation struct {
    Iterations int
}

// RunSimulation performs the Monte Carlo simulation.
func RunSimulation(simulation *MonteCarloSimulation) (float64, error) {
    inside := 0
    for i := 0; i < simulation.Iterations; i++ {
        x, err := rand.Int(rand.Reader, big.NewInt(1000))
        if err != nil {
            return 0, err
        }
        y, err := rand.Int(rand.Reader, big.NewInt(1000))
        if err != nil {
            return 0, err
        }
        if x.Int64()*x.Int64()/1000000 + y.Int64()*y.Int64()/1000000 <= 1 {
            inside++
        }
    }
    return float64(inside) / float64(simulation.Iterations) * 4, nil // pi approximation
}

// Handler for the Monte Carlo simulation.
func MonteCarloHandler(c echo.Context) error {
    params := new(MonteCarloSimulation)
    if err := c.Bind(params); err != nil {
        return err
    }
    if params.Iterations <= 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Iterations must be a positive number")
    }
    pi, err := RunSimulation(params)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "pi": pi,
        "iterations": params.Iterations,
    })
}

func main() {
    e := echo.New()
    e.GET("/simulate", MonteCarloHandler)
    
    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
