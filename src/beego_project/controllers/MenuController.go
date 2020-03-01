package controllers

import (
	"fmt"
	models2 "github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego/orm"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {

	c.LayoutSections = make(map[string]string)
	//{{.footerjs}}
	c.LayoutSections["footerjs"] = "menu/footerjs.html"

	// {{.LayoutContent}}
	//c.setTpl("menu/index.html")
	c.setTpl()
}
func (c *MenuController) List() {
	data, total := models2.MenuList()
	type MenuEx struct {
		models2.MenuModel
		ParentName string
	}
	var menu = make(map[int]string)
	if len(data) > 0 {
		for _, v := range data {
			menu[v.Mid] = v.Name
		}
	}
	var dataEX []MenuEx
	for _, v := range data {
		dataEX = append(dataEX, MenuEx{*v, menu[v.Parent]})
	}
	fmt.Println(total)
	c.Data["LayoutContent"] = dataEX
	//c.listJsonResult(consts.JRCodeSucc, "ok", total, dataEX)

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTpl()
}
func (c *MenuController) Add() {
	//绑定下拉列表
	var pMenus []models2.MenuModel
	data, _ := models2.MenuList()
	for _, v := range data {
		if v.Parent == 0 {
			pMenus = append(pMenus, *v)
		}
	}
	c.Data["PMenus"] = pMenus
	//设置layout_add.html页面的footerjs
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_add.html"
	c.setTpl("menu/add.html", "common/layout_edit.html")
}

func (c *MenuController) AddDo() {
	var m models2.MenuModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Insert(&m)
	}
}

func (c *MenuController) Edit() {
	//获取参数
	c.Data["Mid"] = c.GetString("mid")
	c.Data["Parent"], _ = c.GetInt("parent")
	c.Data["Seq"] = c.GetString("seq")
	c.Data["Name"] = c.GetString("name")

	//绑定下拉列表
	var pMenus []models2.MenuModel
	data, _ := models2.MenuList()
	for _, v := range data {
		if v.Parent == 0 {
			pMenus = append(pMenus, *v)
		}
	}
	c.Data["PMenus"] = pMenus

	//设置layout_edit.html页面的footerjs
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("menu/edit.html", "common/layout_edit.html")
}

func (c *MenuController) EditDo() {
	var m models2.MenuModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Update(&m)
	}
}
