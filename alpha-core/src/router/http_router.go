package router

import (
	"alpha-core/src/config"
	"alpha-core/src/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(r *gin.Engine) {
	r.Use(Cors())
	v1 := r.Group("/alpha/v1")
	{
		v1.GET(config.GetB, controller.GetDataB)
	}

	/***
	1 图像检测
	2 图像识别
	3 视频识别
	4 模型训练
	5 API开放
	**/
	image := r.Group("/alpha/image")
	{
		image.POST(config.Image_identify_path, controller.ExecmagesIdentify)
		image.POST(config.Image_detect_path,controller.ExecImagesDetect)
		image.POST(config.Image_upload_path,controller.UploadImages)

	}
	/***

	1 简介
	2 开发者社区
	3 FAQ
	*/
	about := r.Group("/alpha/about")
	{
		about.POST(config.About_add_message_path, controller.AddMessage)
		about.GET(config.About_query_message_path, controller.QueryAllMessage)
	}
	/**
	下载
	 */
	resource := r.Group("/alpha/resource")
	{
		resource.GET(config.Resource_download_path, controller.FileDownload)
	}
}

// 处理跨域请求,支持OPTIONS方法访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)

		}

		// 处理请求
		c.Next()
	}
}
