package controller

import (
	"alpha-core/src/common"
	"alpha-core/src/config"
	"alpha-core/src/services"
	"alpha-core/src/views"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
图像相关API
 */

func UploadImages(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	saveName := common.GetUUID()
	err = c.SaveUploadedFile(file, config.FILE_UPLOAD_PATH+saveName)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	// 返回文件名以及对应的MD5文件名
	c.JSON(http.StatusOK, gin.H{
		"save_name": saveName,
		"name":      file.Filename,
	})
}

func ExecmagesIdentify(c *gin.Context) {
	var files []views.Fileitem
	if err := c.BindJSON(&files); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var resps []views.ImageIdentifyResp
	ch := make(chan views.ImageIdentifyResp)
	for _, file := range files {
		go services.ImageIdentify(file.FileName, file.FileSaveName, ch)
		// 对结果翻译
		// http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=lynx
		// 并行请求 https://segmentfault.com/a/1190000005064535

	}
	for i := 0; i < len(files); i++ {
		resps = append(resps, <-ch)
	}
	close(ch)
	c.JSON(200, resps)
}

// 图片探测，返回多个图片地址
func ExecImagesDetect(c *gin.Context) {
	var req []views.ImageDetectRequest
	resp := views.ImageDetectResp{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch := make(chan views.ImageDetectItem)
	for _, imageDetect := range req {
		go services.ImageDetect(imageDetect.FileName, imageDetect.FileSaveName, ch)
	}

	var imageDetectItem []views.ImageDetectItem
	for i := 0; i < len(req); i++ {
		//imageDetectItem[i] = <-ch
		imageDetectItem = append(imageDetectItem, <-ch)
	}
	// 关闭通道 不然range会一直下去
	close(ch)
	resp.ImageDetectItems = imageDetectItem
	c.JSON(200, resp)
}
