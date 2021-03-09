package middleware

import (
	"fmt"
	"github.com/ecnuvj/vhoj_api/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				//abort保证被挂起的函数不会被调用，但不会退出当前函数
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		claims, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("auth", claims)
		fmt.Println(claims)
		c.Next()
	}
}

func RoleCheck(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, exist := c.Get("auth")
		if !exist {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		roles := authToken.(*auth.Claims).Roles
		for _, r := range roles {
			if r == role || r == "admin" {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
