package middleware

import (
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"
)



func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        token := session.Get("token")
        if token == nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "GAGAL LOGIN",
            })
            c.Abort()
        } else {
            c.Next()
        }
    }
}


func Session(route *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	route.Use(sessions.Sessions("mysession", store))
}