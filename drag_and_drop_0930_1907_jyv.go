// 代码生成时间: 2025-09-30 19:07:17
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

// DragAndDropApp represents the application structure
type DragAndDropApp struct {
    e *echo.Echo
}

// New initializes the application with necessary middlewares
func New() *DragAndDropApp {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    return &DragAndDropApp{
        e: e,
    }
}

// Start starts the HTTP server
func (app *DragAndDropApp) Start(port string) error {
    // Define routes
    app.e.GET("/", app.homeHandler)

    // Start the server
    return app.e.Start(":" + port)
}

// homeHandler serves the home page with the drag and drop functionality
func (app *DragAndDropApp) homeHandler(c echo.Context) error {
    // Render the HTML template with drag and drop script included
    return c.Render(http.StatusOK, "index", nil)
}

// main function to run the application
func main() {
    app := New()
    port := "8080"
    err := app.Start(port)
    if err != nil {
        // Handle error
        panic(err)
    }
}

// Note: The actual drag and drop functionality will be implemented using JavaScript and HTML.
// This Go code only sets up the Echo server to serve the HTML page.
// The HTML file should include a drag and drop script, such as those available from
// libraries like SortableJS or HTML5 native draggable attributes.
