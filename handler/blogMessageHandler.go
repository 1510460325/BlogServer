package handler

import (
	"blog/dal/domain"
	"blog/info"
	"blog/info/query"
	"blog/service"
	"blog/util"
	"github.com/gin-gonic/gin"
)

func QueryMessage(c *gin.Context, query *query.BlogMessageQuery) (error, *info.PagingInfo) {
	return nil, service.BlogMessageService.QueryPage(query)
}

func CreateMessage(c *gin.Context, record *domain.BlogMessageDomain) (error, bool) {
	err := service.BlogMessageService.Create(record)
	return err, err == nil
}

func DeleteMessage(c *gin.Context, deleteInfo *info.DeleteInfo) (error, bool) {
	ids := util.SplitToIds(*deleteInfo.Ids)
	err := service.BlogMessageService.Delete(ids)
	return err, err == nil
}
