package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateQR(t *testing.T) {
	router := gin.Default()
	router.GET("/create_qr", createQR)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/create_qr?data=test&size=256&level=Medium", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "image/png")
}

func TestCreateBarcode(t *testing.T) {
	router := gin.Default()
	router.GET("/create_barcode", createBarcode)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/create_barcode?data=test&size=200", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "image/png")
}
