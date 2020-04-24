package middleware

import (
	"blog/constant"
	"blog/handler"
	"blog/info"
	"blog/service"
	"blog/util"
	"blog/util/logUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	cookie, err := c.Cookie(constant.SESSION)
	if err != nil || !handler.CheckAccess(cookie) {
		c.JSON(http.StatusInternalServerError, info.JsonResult{
			Code:    constant.FAILED,
			Message: util.StringPtr("no auth"),
		})
		c.Abort()
	}
}

func LogMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logUtil.Logger.Error("error : %v", err)
			c.JSON(http.StatusInternalServerError, info.JsonResult{
				Code:    constant.FAILED,
				Message: util.StringPtr(fmt.Sprintf("%v", err)),
			})
			c.Abort()
		}
	}()
	c.Next()
}

func CorsMiddleware(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func SessionMiddleware(c *gin.Context) {
	_, err := c.Cookie(constant.SESSION)
	if err != nil {
		_ = service.BlogAboutService.AddViewNum()
		uuidV4 := uuid.NewV4().String()
		c.SetCookie(constant.SESSION, uuidV4, constant.MaxAge, "/", "localhost", false, true)
		c.SetCookie(constant.SESSION, uuidV4, constant.MaxAge, "/", "wangzhengyu.cn", false, true)
	}
}
