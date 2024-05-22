package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"image/png"
	"net/http"
	"strconv"
)

func createQR(c *gin.Context) {
	data := c.Query("data")
	sizeStr := c.DefaultQuery("size", "256")
	levelStr := c.DefaultQuery("level", "Medium")

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Size must be an integer"})
		return
	}

	var qrLevel qrcode.RecoveryLevel
	switch levelStr {
	case "Low":
		qrLevel = qrcode.Low
	case "Medium":
		qrLevel = qrcode.Medium
	case "High":
		qrLevel = qrcode.High
	case "Highest":
		qrLevel = qrcode.Highest
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid level. Must be one of: Low, Medium, High, Highest"})
		return
	}

	qrCode, err := qrcode.New(data, qrLevel)
	if err != nil {
		c.Error(err)
		return
	}

	pngImage := qrCode.Image(size)

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, pngImage)
	if err != nil {
		c.Error(err)
		return
	}

	c.Data(http.StatusOK, "image/png", buffer.Bytes())
}
