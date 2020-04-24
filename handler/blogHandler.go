package handler

import (
	"blog/constant"
	"blog/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"

	"blog/dal/domain"
	"blog/info"
	"blog/info/query"
	"blog/service"
)

func QueryBlog(c *gin.Context, query *query.BlogQuery) (error, *info.PagingInfo) {
	return nil, service.BlogService.QueryPage(query)
}

func UpdateBlog(c *gin.Context, record *domain.BlogDomain) (error, bool) {
	if record.ID == nil {
		return errors.New("ID is nil"), false
	}
	err := service.BlogService.Update(record)
	return err, err == nil
}

func CreateBlog(c *gin.Context, record *domain.BlogDomain) (error, bool) {
	err := service.BlogService.Create(record)
	return err, err == nil
}

func DeleteBlog(c *gin.Context, deleteInfo *info.DeleteInfo) (error, bool) {
	ids := util.SplitToIds(*deleteInfo.Ids)
	err := service.BlogService.Delete(ids)
	return err, err == nil
}

func AddViewNum(c *gin.Context, record *domain.BlogDomain) (error, bool) {
	if record.ID == nil {
		return errors.New("ID is nil"), false
	}
	return nil, service.BlogService.AddViewNum(*record.ID)
}

func QueryAbout(c *gin.Context, pageQuery *query.PageQuery) (error, *domain.BlogAboutDomain) {
	return nil, service.BlogAboutService.Query()
}

func UpdateAbout(c *gin.Context, record *domain.BlogAboutDomain) (error, bool) {
	err := service.BlogAboutService.Update(record)
	return err, err == nil
}

func Upload(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	if _, err = os.Open("static"); err != nil {
		if err = os.Mkdir("static", os.ModeDir); err != nil {
			panic(err)
		}
	}
	dst := fmt.Sprintf("static/%v-%v", time.Now().Unix(), header.Filename)
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, info.JsonResult{
		Code: constant.SUCCESS,
		Data: map[string]string{
			"url":      dst,
			"fileName": header.Filename,
		},
	})
}
