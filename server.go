package main

import (
	"blog/config"
	"blog/constant"
	"blog/handler"
	"blog/info"
	"blog/middleware"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

func injectParam(handler interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		fv := reflect.ValueOf(handler)
		modelType := fv.Type().In(1).Elem()
		// 指针类型
		modelPtr := reflect.New(modelType).Interface()
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			err := c.BindJSON(modelPtr)
			if err != nil {
				c.JSON(http.StatusInternalServerError, info.JsonResult{
					Code:    constant.FAILED,
					Message: util.StringPtr(err.Error()),
				})
				return
			}
		} else {
			modelPtr = buildStruct(c, modelType, true)
		}
		// 只有两个参数: (gin.Context, model)
		params := make([]reflect.Value, 2)
		params[0] = reflect.ValueOf(c)
		params[1] = reflect.ValueOf(modelPtr)
		callResult := fv.Call(params)
		if callResult[0].Interface() != nil {
			c.JSON(http.StatusInternalServerError, info.JsonResult{
				Code:    constant.FAILED,
				Message: util.StringPtr(callResult[0].Interface().(error).Error()),
			})
		} else {
			c.JSON(http.StatusOK, info.JsonResult{
				Code: constant.SUCCESS,
				Data: callResult[1].Interface(),
			})
		}
	}
}

func buildStruct(c *gin.Context, modelType reflect.Type, returnPtr bool) interface{} {
	// 指针类型
	modelPtr := reflect.New(modelType).Interface()
	// 真实类型
	modelValue := reflect.ValueOf(modelPtr).Elem()
	num := modelValue.NumField()
	for i := 0; i < num; i++ {
		tagName := modelType.Field(i).Tag.Get("json")
		if modelType.Field(i).Type.Kind() == reflect.Struct {
			modelValue.Field(i).Set(reflect.ValueOf(buildStruct(c, modelType.Field(i).Type, false)))
		} else {
			if queryValue, ok := c.GetQuery(tagName); ok {
				switch modelType.Field(i).Type.Elem().Kind() {
				case reflect.String:
					modelValue.Field(i).Set(reflect.ValueOf(util.StringPtr(queryValue)))
				case reflect.Uint32:
					if value, err := strconv.Atoi(queryValue); err == nil {
						modelValue.Field(i).Set(reflect.ValueOf(util.Uint32Ptr(uint32(value))))
					}
				case reflect.Uint64:
					if value, err := strconv.Atoi(queryValue); err == nil {
						modelValue.Field(i).Set(reflect.ValueOf(util.Uint64Ptr(uint64(value))))
					}
				}
			}
		}
	}
	if returnPtr {
		return modelPtr
	} else {
		return modelValue.Interface()
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LogMiddleware, middleware.SessionMiddleware)
	globalGroup := r.Group("/blog")
	authGroup := globalGroup.Group("/admin")
	authGroup.Use(middleware.AuthMiddleware)
	authGroup.POST("/blog/create", injectParam(handler.CreateBlog))
	authGroup.DELETE("/blog/delete", injectParam(handler.DeleteBlog))
	authGroup.PUT("/blog/update", injectParam(handler.UpdateBlog))
	authGroup.PUT("/about/update", injectParam(handler.UpdateAbout))
	authGroup.POST("/tags/create", injectParam(handler.CreateBlogTags))
	authGroup.DELETE("/tags/delete", injectParam(handler.DeleteBlogTags))
	authGroup.PUT("/tags/update", injectParam(handler.UpdateBlogTags))
	authGroup.DELETE("/message/delete", injectParam(handler.DeleteMessage))
	authGroup.POST("/file/upload", handler.Upload)

	openGroup := globalGroup.Group("/open")
	openGroup.GET("/blog/query", injectParam(handler.QueryBlog))
	openGroup.PUT("/blog/addViewNum", injectParam(handler.AddViewNum))
	openGroup.GET("/tags/query", injectParam(handler.QueryTags))
	openGroup.GET("/message/query", injectParam(handler.QueryMessage))
	openGroup.POST("/message/create", injectParam(handler.CreateMessage))
	openGroup.GET("/about/query", injectParam(handler.QueryAbout))
	openGroup.GET("/about/checkAdmin", injectParam(handler.CheckAdmin))

	globalGroup.StaticFS("/static", http.Dir("static"))
	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()
	fmt.Println("server is started on port:", config.AppSetting.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", config.AppSetting.Server.Port)); err != nil {
		fmt.Println()
		panic("server is err:" + err.Error())
	}
	fmt.Println("server is end!")
}
