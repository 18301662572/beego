package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//demo
//beego框架的orm使用
//导入orm包：github.com/astaxie/beego/orm

type Page struct {
	Id      int64  `form:"id" json:"id"`
	Website string `form:"website" json:"website"`
	Email   string `form:"email" json:"email"`
}

//func init() {
//RegisterDataBase("别名"，"mysql","连接数据库")
//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
//orm.RegisterModel(new(Page))
//}

func GetPage(id int64) *Page {
	//return &Page{
	//	Website: "beego.me",
	//	Email:   "model@beego.com",
	//}

	o := orm.NewOrm()
	p := Page{Id: id}
	err := o.Read(&p)
	if err != nil {
		fmt.Println(err)
	}
	return &p
}

func InserPage() (int64, error) {
	o := orm.NewOrm()
	p := Page{Website: "www.baidu.com", Email: "aa@163.com"}
	id, err := o.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, err)
	return id, err
}

func UpdatePage() (id int64, err error) {
	p := Page{Id: 1, Email: "dd.com"}
	o := orm.NewOrm()
	//Update(model,"修改的字段名")
	id, err = o.Update(&p, "Email")
	return
}
