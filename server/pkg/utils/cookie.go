package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookies(c *gin.Context, token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name: "auth_token",
		Value: token,
		Path: "/",
		Domain: "",
		Expires: time.Now().Add(1*time.Hour),
		MaxAge: 3600,
		Secure: false, //true for production
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Partitioned: true,
	})
}