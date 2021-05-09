package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
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
	serveMux.HandleFunc("/kill", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("is kill"))
		outChan <- struct{}{}
	})

	group.Go(func() error {
		//return server.ListenAndServe()
		//2.设置监听的TCP地址并启动服务
		//参数1:TCP地址(IP+Port)
		//参数2:handler 创建新的*serveMux,不使用默认的
		serveErr := http.ListenAndServe("127.0.0.1:9000", serveMux)
		if serveErr != nil {
			fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", serveErr)
			cancel()
		}

		return nil
	})

	group.Go(func() error {
		// 休眠1秒，用于捕获子协程2的出错
		time.Sleep(1 * time.Second)
		//检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
		err := CheckGoroutineErr(errCtx)
		if err != nil {
			return err
		}
		return nil
	})

	for index := 0; index < 3; index++ {
		indexTemp := index // 子协程中若直接访问index，则可能是同一个变量，所以要用临时变量

		// 新建子协程
		group.Go(func() error {
			if indexTemp == 0 {
				fmt.Println("indexTemp == 0 start ")
				fmt.Println("indexTemp == 0 end")
			} else if indexTemp == 1 {
				fmt.Println("indexTemp == 1 start")
				//这里一般都是某个协程发生异常之后，调用cancel()
				//这样别的协程就可以通过errCtx获取到err信息，以便决定是否需要取消后续操作
				fmt.Println("这里出错了")
				//if index%1 == 0 {
				//	return fmt.Errorf("something has failed on grouting:%d", index)
				//}
				cancel()
				fmt.Println("indexTemp == 1 err ")
			} else if indexTemp == 2 {
				fmt.Println("indexTemp == 2 start")

				// 休眠1秒，用于捕获子协程2的出错
				time.Sleep(1 * time.Second)

				//检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
				err := CheckGoroutineErr(errCtx)
				if err != nil {
					return err
				}
				fmt.Println("indexTemp == 2 end ")
			}
			return nil
		})
	}

	// 捕获err
	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Printf("get error:%v", err)
	}
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


