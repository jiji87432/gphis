package main

import (
	_ "gphis/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
)

func init() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Listen.HTTPPort = 10160

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:prj4gphis@qcloud2017@tcp(123.206.208.96:12321)/gphis?charset=utf8")
}

func main() {
	beego.Run()
}
