package controllers

import (
	"fmt"
	"github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	//赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	//TODO保存用户数据
	fmt.Println("beego:Prepare: " + c.controllerName + "," + c.actionName)
	//权限控制
	user := c.auth()
	c.Data["Menu"] = models.MenuTreeStruct(user)
}

func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要“Controller” 这10个字母
		ctlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctlName + "/" + actionName + ".html"
	}

	_, found := c.Data["Footer"]
	if !found {
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}

//func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int64) {
//	r := &models.ListJsonResult{code, msg, count, obj}
//	c.Data["json"] = r
//	c.ServeJSON()
//	c.StopRun()
//}

//func (c *BaseController) jsonResult(code consts.JRCodeSucc, msg string, obj interface{}) {
//	r := &models.JsonResult{code, msg, obj}
//	c.Data["json"] = r
//	c.ServeJSON()
//	c.StopRun()
//}

func (c *BaseController) auth() models.UseModel {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect("/login", 302)
		c.StopRun()
		return models.UseModel{}
	} else {
		return user.(models.UseModel)
	}
}
