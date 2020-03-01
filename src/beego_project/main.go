package main

import (
	_ "github.com/18301662572/beego/src/beego_project/routers"
	_ "github.com/18301662572/beego/src/beego_project/sysint"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//导入日志包 ：github.com/astaxie/beego/logs

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}
