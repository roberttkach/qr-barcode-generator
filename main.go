package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)
	log.SetOutput(logFile)

	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(errorHandlingMiddleware())

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/create_qr", createQR)
	router.GET("/create_barcode", createBarcode)

	err = router.Run(":" + config.Port)
	if err != nil {
		return
	}
}
