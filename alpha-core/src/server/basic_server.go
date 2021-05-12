package server

import (
	"alpha-core/src/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//TODO: 指针函数还需要细看   https://www.cnblogs.com/justdoyou/p/10026004.html
func Get_basic_server(r *gin.Engine) *http.Server {
	h := (http.Handler)(r)
	s := &http.Server{
		Addr:           config.PORT,
		Handler:        h,
		ReadTimeout:    config.READ_TIME_OUT * time.Second,
		WriteTimeout:   config.WRITE_TIME_OUT * time.Second,
		MaxHeaderBytes: config.MAX_HEADER_BYTE,
	}
	return s
}

// 另外还可以实现多个server
