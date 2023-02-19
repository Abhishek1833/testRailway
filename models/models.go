package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Item struct {
	Id          int    `json:"id"; orm:"column(id);auto"`
	Description string `json:"description"; orm:"column(description)"`
	Type        string `json:"type"; orm:"column(type)"`
	Name        string `json:"name"; orm:"column(name)"`
}

func (t *Item) TableName() string {
	return "item"
}
func init() {
	orm.RegisterModel(new(Item))
}
func PostItem(pst *Item) (id int64, err error) {
	fmt.Println("model")
	o := orm.NewOrm()
	id, err = o.Insert(pst)
	return id, err
}

func GetAllItems() (v []Item, err error) {
	o := orm.NewOrm()
	v = []Item{}
	_, err = o.QueryTable(new(Item)).RelatedSel().All(&v)
	return v, err
}
func GetItems(id int) (v *Item, err error) {
	o := orm.NewOrm()
	v = &Item{}
	err = o.QueryTable(new(Item)).Filter("id", id).RelatedSel().One(v)
	return v, err
}

func UpdateItem(m *Item) (err error) {
	o := orm.NewOrm()
	v := Item{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteItem(id int) (err error) {
	o := orm.NewOrm()
	v := Item{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Item{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
