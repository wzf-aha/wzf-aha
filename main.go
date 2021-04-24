package main

import (
	"fmt"
	"ginDemo/app/student"
	"ginDemo/app/teacher"
	"ginDemo/common"
	"ginDemo/routers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/trace"
	"io"
	"net/http"
	"os"
)

func main()  {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//go startTrace()

	common.SysTeamInfo()
	//加载多个APP路由配置
	routers.Include(student.Routers,teacher.Routers)
	//初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	//grpclog.Println("Trace listen on 50051")
	fmt.Println("Trace listen on 50051")
}

