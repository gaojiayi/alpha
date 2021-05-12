package controller

import (
	"alpha-core/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
资源相关API
下载
*/

func FileDownload(c *gin.Context) {
	filename := c.Param("fileName")
	// fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(config.FILE_DOWNLOAD_PATH + filename)
}
