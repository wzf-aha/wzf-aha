package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	//1.注册一个处理器函数
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("start server"))
	})

	//传入一个退出信号，该信号待被接收处理
	outChan := make(chan struct{})
	serveMux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("stop server"))
		outChan <- struct{}{}
	})

	server := http.Server{
		Handler: serveMux,
		Addr:    "127.0.0.1:8080",
	}

	//监听HTTPSERVER
	group.Go(func() error {
		return server.ListenAndServe()
	})
	//2.设置监听的TCP地址并启动服务
	//参数1:TCP地址(IP+Port)
	//参数2:handler 创建新的*serveMux,不使用默认的
	//serveErr := http.ListenAndServe("127.0.0.1:9000",serveMux)
	//if serveErr != nil {
	//	fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", serveErr)
	//	cancel()
	//}

	//监听错误及退出信号
	group.Go(func() error {
		select {
		case <-errCtx.Done():
			fmt.Println("errgroup exit")
		case <-outChan:
			fmt.Println("server will stop")
		}
		cancel()
		fmt.Println("will stop")

		return server.Shutdown(errCtx)
	})

	//signal信号监听 如何注册？
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM,syscall.SIGUSR2,syscall.SIGQUIT)
	group.Go(func() error {
		select {
		case <-errCtx.Done():
			return errCtx.Err()
		case sig := <-ch:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	// 捕获err
	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Println("get error:%v", err)
	}
	cancel()
}

//校验是否有协程已发生错误
func CheckGoroutineErr(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}


