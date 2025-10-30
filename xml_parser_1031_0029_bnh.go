// 代码生成时间: 2025-10-31 00:29:20
package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/labstack/echo"
)

// XMLData represents the structure of the XML data we expect to parse.
type XMLData struct {
    Items []struct {
        Name  string `xml:"name,attr"`
        Value string `xml:"value"`
    } `xml:"item"`
}

// parseXML parses the XML data from a string and returns an XMLData struct.
func parseXML(data string) (*XMLData, error) {
    var result XMLData
    err := xml.Unmarshal([]byte(data), &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// XMLParser is a handler function that takes an XML string, parses it, and
// returns the parsed XML data as a JSON response.
func XMLParser(c echo.Context) error {
    // Read the XML data from the request body.
    xmlData, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        return err
    }
    defer c.Request().Body.Close()

    // Parse the XML data.
    parsedData, err := parseXML(string(xmlData))
    if err != nil {
        return err
    }

    // Return the parsed data as JSON.
    return c.JSON(http.StatusOK, parsedData)
}

func main() {
    e := echo.New()
    e.POST("/parse", XMLParser)
    e.Logger.Fatal(e.Start(":8080"))
}
