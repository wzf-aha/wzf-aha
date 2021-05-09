package controller

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/pkg/errors"
)

func worktwo()  {

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
	//fmt.Printf("original error: %T %v\n",errors.Cause(err),errors.Cause(err))
	//fmt.Printf("stack trace:\n%+v\n",err)

	if err != nil {
		//记录堆栈详情
		fmt.Printf("errinfo:%+v\n",err)

		if errors.Is(err,sql.ErrNoRows){//对于未获取到结果集的返回值做对应的业务处理
			fmt.Print("哦，没找到匹配的数据呀")
		}else{// 当有非sql.ErrNoRows的错误产生时，则报错
			fmt.Print("出错了")
		}
	}

	fmt.Print(year)
}

/**
 *  在获取单行数据时，Scan没有匹配到结果集时会报sql.ErrNoRows错误
 *	dao 层中当遇到一个 sql.ErrNoRows 的时候，不应该Wrap 这个 error，抛给上层
 *  dao 层只负责查询或操作数据，不对结果做处理，错误和结果应由上层的业务逻辑来判断是否要处理及怎么处理
 *	dao 层具有比较高的重用性，最好返回根错误供业务侧判别
 */
func TestSelect(db *sql.DB) (year int,err error) {
	err = db.QueryRow("select year from books where id = ?", 1).Scan(&year)
	//if err != nil {//添加上能帮助定位问题或者问题数据的信息往上抛
	//	return 0, errors.Wrap(err,"err id:1")
	//}
	return
}