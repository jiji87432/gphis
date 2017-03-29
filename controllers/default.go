package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (this *DefaultController) Get() {
	this.Data["json"] = &BaseResp{Result: true, Reason: "Graduate Project - Hospital Information System"}
	this.ServeJSON()
}
