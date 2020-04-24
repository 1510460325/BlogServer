package handler

import (
	"blog/constant"
	"blog/info"
	"blog/service"
	"blog/util"
	"errors"
	"github.com/bluele/gcache"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	sessionCache = gcache.New(1000000).LRU().Build()
	captchaStore = base64Captcha.NewMemoryStore(1000000, 5*time.Minute)
	audioDriver  = base64Captcha.NewDriverAudio(4, "zh")
	stringDriver = base64Captcha.NewDriverMath(35, 100, 0, base64Captcha.OptionShowSineLine, nil, nil)
	audioCaptcha  = &base64Captcha.Captcha{}
	stringCaptcha = &base64Captcha.Captcha{}
)

func init() {
	audioCaptcha = base64Captcha.NewCaptcha(audioDriver, captchaStore)
	stringCaptcha = base64Captcha.NewCaptcha(stringDriver, captchaStore)
}

func CheckAccess(session string) bool {
	_, err := sessionCache.Get(session)
	return err == nil
}

func CheckAdmin(c *gin.Context, adminInfo *info.CheckAdminInfo) (error, bool) {
	if adminInfo.Pwd == nil {
		return errors.New("pwd is nil"), false
	}
	access := service.BlogAboutService.CheckAdmin(*adminInfo.Pwd)
	if access {
		cookies, err := c.Cookie(constant.SESSION)
		if err != nil {
			uuidV4 := uuid.NewV4().String()
			c.SetCookie(constant.SESSION, uuidV4, constant.MaxAge, "/", "localhost", false, true)
			c.SetCookie(constant.SESSION, uuidV4, constant.MaxAge, "/", "wangzhengyu.cn", false, true)
			cookies = uuidV4
		}
		err = sessionCache.SetWithExpire(cookies, true, constant.MaxAge*time.Second)
	}
	return nil, access
}

func GetCaptcha(c *gin.Context, query *info.CaptchaInfo) (error, *info.CaptchaInfo) {
	if query.IsAudio != nil && *query.IsAudio {
		id, b64s, err := audioCaptcha.Generate()
		return err, &info.CaptchaInfo{
			ID:         util.StringPtr(id),
			Base64Data: util.StringPtr(b64s),
		}
	} else {
		id, b64s, err := stringCaptcha.Generate()
		return err, &info.CaptchaInfo{
			ID:         util.StringPtr(id),
			Base64Data: util.StringPtr(b64s),
		}
	}
}