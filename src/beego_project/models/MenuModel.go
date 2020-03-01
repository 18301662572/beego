package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"sort"
)

//菜单模块

//gorm中文文档： http://gorm.book.jasperxu.com/

//MenuModel 菜单模块
type MenuModel struct {
	Mid    int    `orm:"pk;auto;column(Mid)"` //orm:表示字段约束
	Parent int    `orm:"column(Parent)"`
	Name   string `orm:"size(45);column(Name)"`
	Seq    int    `orm:"column(Seq)"`
	Format string `orm:"size(2048);default({});column(Format)"`
}

//MenuTree 菜单树
type MenuTree struct {
	MenuModel
	Child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "cms_menu"
}

func MenuStruct() map[int]MenuTree {
	query := orm.NewOrm().QueryTable("cms_menu")
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	var menu = make(map[int]MenuTree)
	if len(data) > 0 {
		for _, v := range data {
			if v.Parent == 0 {
				var tree = new(MenuTree)
				tree.MenuModel = *v
				menu[v.Mid] = *tree
			} else {
				if tmp, ok := menu[v.Parent]; ok {
					tmp.Child = append(tmp.Child, *v)
					menu[v.Parent] = tmp
				}
			}
		}
	}
	return menu
}

func MenuTreeStruct(user UseModel) map[int]MenuTree {
	query := orm.NewOrm().QueryTable("cms_menu")
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	var menu = make(map[int]MenuTree)
	//auth
	if len(user.AuthStr) > 0 {
		var authArr []int
		json.Unmarshal([]byte(user.AuthStr), &authArr)
		sort.Ints(authArr)

		for _, v := range data { //查询出来的数组
			//fmt.Println(v.Mid, v.Parent, v.Name)
			if v.Parent == 0 {
				idx := sort.SearchInts(authArr, v.Mid)
				found := (idx < len(authArr) && authArr[idx] == v.Mid)
				if found {
					var tree = new(MenuTree)
					tree.MenuModel = *v
					menu[v.Mid] = *tree
				}
			} else {
				if tmp, ok := menu[v.Parent]; ok {
					tmp.Child = append(tmp.Child, *v)
					menu[v.Parent] = tmp
				}
			}
		}
	}
	return menu
}

func MenuList() ([]*MenuModel, int64) {
	query := orm.NewOrm().QueryTable("cms_menu")
	total, _ := query.Count()
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	return data, total
}

//ParentMenuList 获取一级节点
func ParentMenuList() []*MenuModel {
	query := orm.NewOrm().QueryTable("cms_menu").Filter("parent", 0)
	data := make([]*MenuModel, 0)
	query.OrderBy("-seq").All(&data)
	return data
}
func MenuFormatStruct(mid int) *simplejson.Json {
	menu := MenuModel{Mid: mid}
	err := orm.NewOrm().Read(&menu)
	if err == nil {
		jsonstruct, err2 := simplejson.NewJson([]byte(menu.Format))
		if err2 == nil {
			return jsonstruct
		}
	}
	return nil
}
