package controllers

import (
	"fmt"
	"github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//储存session
	//储存session之前要在conf->app.conf里面配置sessionon=true
	c.SetSession("user", "Kelly")
	//获取session
	user := c.GetSession("user")
	fmt.Println(user)

	//记录日志
	logs.Informational("user Kelly loged in")

	//添加数据
	//id, err := models.InserPage()
	//if err != nil {
	//	fmt.Println("insert page err:", err)
	//	return
	//}
	//修改数据
	//id, _ := models.UpdatePage()
	//查看数据
	page := models.GetPage(1)
	c.Data["Website"] = page.Website
	c.Data["Email"] = page.Email
	c.TplName = "index.tpl"

}
