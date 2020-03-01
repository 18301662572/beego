package controllers

import (
	"fmt"
	"github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type FormatController struct {
	BaseController
}

//渲染页面
func (c *FormatController) Edit() {
	midvalue, _ := c.GetInt("mid")
	menu := models.MenuModel{Mid: midvalue}
	orm.NewOrm().Read(&menu)

	c.Data["Mid"] = midvalue
	c.Data["Format"] = menu.Format

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "format/footerjs_edit.html"
	c.setTpl("format/edit.html", "common/layout_edit.html")
}

//修改提交页面
func (c *FormatController) EditDo() {
	mid, _ := c.GetInt("mid")
	f := c.GetString("formatstr") //提交的format信息
	if 0 != mid {
		menu := models.MenuModel{Mid: mid, Format: f}
		mid, _ := orm.NewOrm().Update(&menu, "format")
		fmt.Println(mid)
		//刷新当前页面
		c.Redirect("/format/editdo?mid="+strconv.FormatInt(mid, 10), 302)
		//c.jsonResult(consts.JRCodeSucc,"ok",mid)
		c.StopRun()
	}
	//c.jsonResult(consts.JRCodeFailed, "", 0)
}
