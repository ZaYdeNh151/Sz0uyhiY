// 代码生成时间: 2025-10-28 18:20:10
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "net/http"
    "os"
    "path/filepath"

    "github.com/disintegration/imaging"
    "github.com/labstack/echo"
)

// ImageFilterProcessor 定义图像处理函数的类型
type ImageFilterProcessor func(image.Image) image.Image

// ImageFilterEngine 结构体，包含滤镜列表
type ImageFilterEngine struct {
    Filters map[string]ImageFilterProcessor
}

// NewImageFilterEngine 创建一个新的ImageFilterEngine实例
func NewImageFilterEngine() *ImageFilterEngine {
    return &ImageFilterEngine{
        Filters: make(map[string]ImageFilterProcessor),
    }
}

// RegisterFilter 注册一个新的滤镜处理器
func (engine *ImageFilterEngine) RegisterFilter(name string, filter ImageFilterProcessor) {
    engine.Filters[name] = filter
}

// ApplyFilter 应用滤镜到图像
func (engine *ImageFilterEngine) ApplyFilter(name string, img image.Image) (image.Image, error) {
    if filter, exists := engine.Filters[name]; exists {
        return filter(img), nil
    }
    return nil, fmt.Errorf("filter '%s' not found", name)
}

// HandleImageFilter 处理图像滤镜的HTTP请求
func HandleImageFilter(engine *ImageFilterEngine, c echo.Context) error {
    // 从请求中获取滤镜名称和图像文件名
    filterName := c.QueryParam("filter")
    file, err := c.FormFile("image")
    if err != nil {
        return err
    }
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    // 读取图像
    img, _, err := image.Decode(src)
    if err != nil {
        return err
    }

    // 应用滤镜
    filteredImg, err := engine.ApplyFilter(filterName, img)
    if err != nil {
        return err
    }

    // 保存滤镜后的图像
    hash := fmt.Sprintf("%x", md5.Sum([]byte(file.Filename)))
    dstPath := filepath.Join("./filtered", hash+".jpg")
    dst, err := os.Create(dstPath)
    if err != nil {
        return err
    }
    defer dst.Close()

    if err := jpeg.Encode(dst, filteredImg, nil); err != nil {
        return err
    }

    // 返回保存的图像路径
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Image processed successfully",
        "path": dstPath,
    })
}

func main() {
    e := echo.New()
    engine := NewImageFilterEngine()

    // 注册一个简单的灰度滤镜
    engine.RegisterFilter("grayscale", imaging.Grayscale)

    // 设置路由和处理器
    e.POST("/filter", HandleImageFilter(engine))

    // 启动服务器
    e.Start(":8080")
}