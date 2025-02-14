package middleware

import (
	"net/http"

	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Lấy token từ header Authorization
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, "Authorization token required")
		c.Abort()
		return
	}

	// Xác thực và giải mã token
	claims, err := token.VerifyJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid or expired token")
		c.Abort()
		return
	}

	// Lấy userID từ claims
	ID, ok := claims["ID"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, "User ID not found in token")
		c.Abort()
		return
	}

	// Lưu userID vào context để các handler khác có thể truy xuất
	c.Set("ID", ID)
	c.Next()
}
