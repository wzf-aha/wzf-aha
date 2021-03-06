package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/trace"
	"log"
	"net/http"
)

func main()  {
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	////gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	////go startTrace()
	//
	//common.SysTeamInfo()
	////加载多个APP路由配置
	//routers.Include(student.Routers,teacher.Routers)
	////初始化路由
	//r := routers.Init()
	//if err := r.Run(); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	year,err := TestSelect(db)

	//1.不允许有未匹配到的情况
	if err != nil {
		fmt.Print("出错了")
	}

	//2.允许有未匹配到的情况
	if err != nil {
		if err != sql.ErrNoRows {// 当有非sql.ErrNoRows的错误产生时，则报错
			fmt.Print("出错了")
		}else{
			//对于未获取到结果集的返回值做对应的业务处理
		}
	}

	fmt.Print(year)
}

/**
 *  在获取单行数据时，Scan没有匹配到结果集时会报sql.ErrNoRows错误
 *	dao 层中当遇到一个 sql.ErrNoRows 的时候，应该Wrap 这个 error，抛给上层
 *  dao 层只负责查询或操作数据，不对结果做处理，错误和结果应由上层的业务逻辑来判断是否要处理及怎么处理
 */
func TestSelect(db *sql.DB) (year int,err error) {
	err = db.QueryRow("select year from books where id = ?", 1).Scan(&year)
	return
}



func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	//grpclog.Println("Trace listen on 50051")
	fmt.Println("Trace listen on 50051")
}

