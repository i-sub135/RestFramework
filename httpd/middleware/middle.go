package middleware

import (
	"net/http"
	"os"
	"rest-framework/httpd/response"

	"github.com/gin-gonic/gin"
)

// RestMiddleware -- Reciver middleware
type RestMiddleware struct {
}

var rest response.RestResponse

// ValidREQ -- Validation Requests
func (r *RestMiddleware) ValidREQ(res *gin.Context) {
	methode := res.Request.Method
	if methode == "GET" {
		api := res.GetHeader("api-accsess")
		access := os.Getenv("APIACCSES")
		if api == "" || api != access {
			res.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Akses ditolak",
				"result":  []map[string]string{},
			})
			res.Abort()
			return
		}
	}
	if methode == "POST" {
		token := res.GetHeader("x-access-token")

		if token == "" {
			res.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Token can not be empty",
				"result":  map[string]string{},
			})
			res.Abort()
			return
		}
		_, err := ValidJwt(token)
		if err != nil {
			res.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Token Not Valid",
				"result": map[string]interface{}{
					"error": err.Error(),
				},
			})
			res.Abort()
		}
		res.Next()
	}
}

// ErrorHandle -- error handler
func (r *RestMiddleware) ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		if status == 404 {
			c.JSON(http.StatusNotFound, rest.RespNotFound(401, "page Not found"))
			c.Abort()
			return
		}
	}
}
