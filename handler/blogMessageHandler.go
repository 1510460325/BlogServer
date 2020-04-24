package handler

import (
	"blog/info"
	"blog/info/query"
	"blog/service"
	"blog/util"
	"errors"
	"github.com/gin-gonic/gin"
)

func QueryMessage(c *gin.Context, query *query.BlogMessageQuery) (error, *info.PagingInfo) {
	return nil, service.BlogMessageService.QueryPage(query)
}

func CreateMessage(c *gin.Context, record *info.CreateMessageInfo) (error, bool) {
	if record.CaptchaID == nil || record.VerifyValue == nil ||
		!captchaStore.Verify(*record.CaptchaID, *record.VerifyValue, true) {
		return errors.New("captcha verify failed"), false
	}
	err := service.BlogMessageService.Create(&record.BlogMessageDomain)
	return err, err == nil
}

func DeleteMessage(c *gin.Context, deleteInfo *info.DeleteInfo) (error, bool) {
	ids := util.SplitToIds(*deleteInfo.Ids)
	err := service.BlogMessageService.Delete(ids)
	return err, err == nil
}
