package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

type DataModel struct {
	Did        int    `orm:"pk;auto;column(Did)"`
	Mid        int    `orm:"default(0);column(Mid)"`
	Parent     int    `orm:"default(0);column(Parent)"`
	Name       string `orm:"size(64);column(Name)"`
	Content    string `orm:"size(2048);default({});column(Content)"`
	Seq        int    `orm:"index;column(Seq)"`
	Status     int8   `orm:"column(Status)"`
	UpdateTime int64  `orm:"column(UpdateTime)"`
}

func (m *DataModel) TableName() string {
	return "cms_data"
}

func DataList(mid, pageSize, page int) ([]map[string]interface{}, int64) {
	if mid <= 0 {
		return nil, 0
	}
	//处理分页信息
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable("cms_data").Filter("mid", mid)
	total, _ := query.Count()
	data := make([]*DataModel, 0)
	query.OrderBy("parent", "-seq").Limit(pageSize, offset).All(&data)
	dataEX := make([]map[string]interface{}, 0)
	for _, v := range data {
		sj, err := simplejson.NewJson([]byte(v.Content))
		if err == nil {
			sj.Set("did", v.Did)
			sj.Set("name", v.Name)
			sj.Set("mid", v.Mid)
			sj.Set("parent", v.Parent)
			sj.Set("seq", v.Seq)
			sj.Set("status", v.Status)
			sj.Set("updatetime", v.UpdateTime)
			dataEX = append(dataEX, sj.MustMap())
		}
	}
	return dataEX, total
}

func DataRead(did int) *simplejson.Json {
	if did <= 0 {
		return nil
	}
	data := DataModel{Did: did}
	err := orm.NewOrm().Read(&data)
	if err == nil {
		sj, err2 := simplejson.NewJson([]byte(data.Content))
		if err2 == nil {
			sj.Set("did", data.Did)
			sj.Set("name", data.Name)
			sj.Set("mid", data.Mid)
			sj.Set("parent", data.Parent)
			sj.Set("seq", data.Seq)
			sj.Set("status", data.Status)
			sj.Set("updatetime", data.UpdateTime)
			return sj
		}
	}
	return nil
}
