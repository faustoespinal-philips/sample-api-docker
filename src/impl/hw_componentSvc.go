package impl

import (
	"github.com/gin-gonic/gin"
)

// GetHwComponent -
func GetHwComponent(c *gin.Context) gin.H {
	return gin.H{
		"device-type": "transducer",
		"id":          "AFDAFEA232-2432AADFAD",
	}
}
