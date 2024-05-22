package main

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"strconv"
)

func createBarcode(c *gin.Context) {
	data := c.Query("data")
	sizeStr := c.DefaultQuery("size", "200")
	
	if data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data query parameter is required"})
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Size must be an integer"})
		return
	}

	barCode, err := code128.Encode(data)
	if err != nil {
		c.Error(err)
		return
	}

	scaled, err := barcode.Scale(barCode, size, size)
	if err != nil {
		c.Error(err)
		return
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, scaled)
	if err != nil {
		c.Error(err)
		return
	}

	c.Data(http.StatusOK, "image/png", buffer.Bytes())
}
