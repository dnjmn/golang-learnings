package main

import "github.com/astaxie/beego/orm"

func main() {
}

type my_table struct {
	id  int    `orm:"column(id)"`
	val string `orm:"column(val)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/?charset=utf8")
}
