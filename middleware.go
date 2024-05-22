package main

import (
	"github.com/gin-gonic/gin"
)

func errorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			c.JSON(c.Writer.Status(), gin.H{"error": err.Error()})
		}
	}
}
