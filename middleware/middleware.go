package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordLogAndTime(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	oldTime := time.Now()
	c.Next()
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("admin_id") == nil && c.Request.URL.Path != "/login" && c.Request.URL.Path != "/health" {
			session.AddFlash("ログインしてください。", "Err")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/login")
		} else {
			c.Next()
		}
	}
}
