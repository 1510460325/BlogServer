package handler

import (
	"blog/constant"
	"blog/util"
	"errors"
	"fmt"
	"github.com/bluele/gcache"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"time"

	"blog/dal/domain"
	"blog/info"
	"blog/info/query"
	"blog/service"
)

var adminSession gcache.Cache = nil

func init() {
	adminSession = gcache.New(1000000).LRU().Build()
}

func CheckAccess(session string) bool {
	_, err := adminSession.Get(session)
	return err == nil
}

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
		err = adminSession.SetWithExpire(cookies, true, constant.MaxAge*time.Second)
	}
	return nil, access
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
