// 代码生成时间: 2025-10-06 20:06:34
package main

import (
    "fmt"
    "net/http"
    "log"
    "time"

    "github.com/labstack/echo"
)

// TrafficData 存储网络流量信息
type TrafficData struct {
    Timestamp time.Time `json:"timestamp"`
    BytesSent  int64     `json: