package controllers

import (
	"github.com/18301662572/beego/src/beego_project/models"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		//验证用户名密码
		userkey := strings.TrimSpace(c.GetString("userkey"))
		pwd := strings.TrimSpace(c.GetString("password"))
		if len(userkey) > 0 && len(pwd) > 0 {
			user := models.GetUserByName(userkey)
			if user.Password == pwd {
				c.SetSession("user", user)
				c.Redirect("/menu", 302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}
