package main

import (
	"appdemo/controller/app"

	"appdemo/gitversion"
	"bdlib/util"
	"fargo"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {

	flag.Parse()
	args := flag.Args()

	// 只显示版本
	if util.InSlice("version", args) {
		fmt.Println(gitversion.Version)
		return
	}

	// 初始化资源
	if err := Init(); err != nil {
		fargo.Error(err)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// controller下的目录对应各个业务模块

	// fargo.Add("/app/test", &app.Controller{}) // 添加固定路由 http://127.0.0.1:6060/app/test
	fargo.Add("/app/:do", &app.Controller{}) // 添加正则路由, 建议使用 http://127.0.0.1:6060/app/xxx

	resetPROCS()

	fargo.Infof("appdemo start ...")
	fargo.Run()
}

// 设置并行数，实例cpu数减一
func resetPROCS() {
	num := runtime.NumCPU()
	if num > 1 {
		num--
	}
	runtime.GOMAXPROCS(num)
}
