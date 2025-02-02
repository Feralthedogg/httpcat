package main

import (
	"httpcat/pkg/httpcat"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		codeParam := c.Query("code")
		if codeParam == "" {
			c.String(http.StatusOK, "Hello from Gin! Try ?code=404")
			return
		}

		code, err := strconv.Atoi(codeParam)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid code param")
			return
		}

		if code >= 400 && code <= 599 {
			httpcat.SendError(c.Writer, code)
			return
		}

		c.String(http.StatusOK, "Non-error code: %d", code)
	})

	r.Run(":8080")
}
