package models

import "github.com/astaxie/beego/orm"

type UseModel struct {
	UserId   int    `orm:"pk;auto;column(UserId)"`
	UserKey  string `orm:"size(64);unqiue;column(UserKey)"`
	UserName string `orm:"size(64);column(UserName)"`
	AuthStr  string `orm:"size(512);column(AuthStr)"` //用户权限
	Password string `orm:"size(128);column(Password)"`
	IsAdmin  int8   `orm:"default(0);column(IsAdmin)"`
}

func (m *UseModel) TableName() string {
	return "cms_user"
}

func (m *UseModel) UserList(pageSize, page int) ([]*UseModel, int64) {
	query := orm.NewOrm().QueryTable("cms_user")
	data := make([]*UseModel, 0)
	offset := (page - 1) * pageSize
	total, _ := query.OrderBy("-user_id").Limit(pageSize, offset).All(&data)
	return data, total
}

func GetUserByName(userkey string) UseModel {
	user := UseModel{UserKey: userkey}
	o := orm.NewOrm()
	o.Read(&user, "UserKey")
	return user
}
