package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)
func MyFMT() gin.HandlerFunc {

	return func(c *gin.Context) {
		host := c.Request.Host
		fmt.Printf("Before: %s\n",host)
		c.Next()
		fmt.Println("Next: ...")
	}

}
