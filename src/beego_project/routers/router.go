package routers

import (
	"github.com/18301662572/beego/src/beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//login
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

	//menu
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/add", &controllers.MenuController{}, "Get:Add")
	beego.Router("/menu/adddo", &controllers.MenuController{}, "*:AddDo")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "Get:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")

	//user
	beego.Router("/user", &controllers.UserController{}, "Get:Index")
	beego.Router("/user/list", &controllers.UserController{}, "*:List")
	beego.Router("/user/add", &controllers.UserController{}, "Get:Add")
	beego.Router("/user/adddo", &controllers.UserController{}, "*:AddDo")
	beego.Router("/user/edit", &controllers.UserController{}, "Get:Edit")
	beego.Router("/user/editdo", &controllers.UserController{}, "*:EditDo")

	//format
	beego.Router("/format/edit", &controllers.FormatController{}, "Get:Edit")
	beego.Router("/format/editdo", &controllers.FormatController{}, "*:EditDo")

	//data
	//访问网址：127.0.0.1:8080/data/1
	//beego.Router("/data/index/?:mid", &controllers.FormatController{}, "Get:Index")
	//beego.Router("/data/list/?:mid", &controllers.FormatController{}, "*:List")
	beego.Router("/data/edit", &controllers.FormatController{}, "Get:Edit")
	beego.Router("/data/editdo", &controllers.FormatController{}, "*:EditDo")

}
