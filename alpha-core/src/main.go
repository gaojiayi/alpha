package main

//  解释了包依赖的概念  https://www.jianshu.com/p/bbed916d16ea
import (
	"alpha-core/src/router"
	"fmt"
	"github.com/gin-gonic/gin"

	"alpha-core/src/server"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// todo 日志的支持
	// Force log's color
	//gin.ForceConsoleColor()
	gin.DisableConsoleColor()

	// Logging to a file. 必须写在上面
	f, _ := os.Create("alpha-core.log")
	// 同时输出控制台以及日志文件
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.New()
	router.SetRouter(r)

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	//r = gin.New() //  注意  router不要使用默认的方式创建
	//router.SetRouter(r)
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	srv := server.Get_basic_server(r)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
