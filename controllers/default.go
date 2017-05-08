package controllers

import (
	"github.com/astaxie/beego"
)

type BaseResp struct {
	Result bool        `json:"result"`
	Reason interface{} `json:"reason"`
}

type DefaultController struct {
	beego.Controller
}

func (this *DefaultController) Get() {
	this.succ("Graduate Project - Hospital Information System")
	this.ServeJSON()
}

func (this *DefaultController) succ(content interface{}) {
	this.Data["json"] = &BaseResp{Result: true, Reason: content}
}
