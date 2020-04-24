package handler

import (
	"errors"
	"github.com/gin-gonic/gin"

	"blog/dal/domain"
	"blog/info"
	"blog/info/query"
	"blog/service"
	"blog/util"
)

func QueryTags(c *gin.Context, query *query.PageQuery) (error, []*domain.BlogTagsDomain) {
	return nil, service.BlogTagsService.Query()
}

func UpdateBlogTags(c *gin.Context, record *domain.BlogTagsDomain) (error, bool) {
	if record.ID == nil {
		return errors.New("ID is nil"), false
	}
	err := service.BlogTagsService.Update(record)
	return err, err == nil
}

func CreateBlogTags(c *gin.Context, record *domain.BlogTagsDomain) (error, bool) {
	err := service.BlogTagsService.Create(record)
	return err, err == nil
}

func DeleteBlogTags(c *gin.Context, deleteInfo *info.DeleteInfo) (error, bool) {
	ids := util.SplitToIds(*deleteInfo.Ids)
	err := service.BlogTagsService.Delete(ids)
	return err, err == nil
}
