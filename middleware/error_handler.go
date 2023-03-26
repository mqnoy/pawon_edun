package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		// log, handle, etc.
		logrus.Errorln(err)
	}

	c.JSON(http.StatusInternalServerError, "")
}
