package controllers

import (
	"encoding/json"
	"github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {

	c.LayoutSections = make(map[string]string)
	//{{.footerjs}}
	c.LayoutSections["footerjs"] = "user/footerjs.html"

	// {{.LayoutContent}}
	//c.setTpl("user/index.html")
	c.setTpl()
}
func (c *UserController) List() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs.html"
	c.setTpl()
}
func (c *UserController) Add() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_add.html"
	c.setTpl("user/add.html", "common/layout_edit.html")
}

func (c *UserController) AddDo() {
	var m models.UseModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Update(&m)
	}
}

//渲染编辑页面
func (c *UserController) Edit() {
	uid, _ := c.GetInt("uid")
	o := orm.NewOrm()
	var user = models.UseModel{UserId: uid}
	o.Read(&user)
	user.Password = ""
	c.Data["User"] = user

	//绑定用户权限
	authmap := make(map[int]bool)
	if len(user.AuthStr) > 0 {
		var authobj []int
		json.Unmarshal([]byte(user.AuthStr), &authobj)
		for _, v := range authobj {
			authmap[v] = true
		}
	}
	type Menuitem struct {
		Name    string
		Ischeck bool
	}
	menu := models.ParentMenuList()
	menus := make(map[int]Menuitem)
	for _, v := range menu {
		menus[v.Mid] = Menuitem{v.Name, authmap[v.Mid]}
	}
	c.Data["Menus"] = menus

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/edit.html", "common/layout_edit.html")

}

//提交编辑页面
func (c *UserController) EditDo() {
	var m models.UseModel
	if err := c.ParseForm(&m); err == nil {
		orm.NewOrm().Update(&m)
	}
}
