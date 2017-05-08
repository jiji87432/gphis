package controllers

import (
	"github.com/astaxie/beego"
	"gphis/models"
)

type DrugController struct {
	beego.Controller
}

func (this *DrugController) Get() {
	switch this.Ctx.Input.Param(":type") {
	case "all":
		this.GetAll()
	default:
		this.notFound()
	}

	this.ServeJSON()
}

func (this *DrugController) GetAll() {
	d := models.T_Drug{}
	ret, err := d.GetAll()
	if err != nil {
		this.fail(err.Error())
		return
	}

	this.succ(ret)
}

func (this *DrugController) succ(content interface{}) {
	this.Data["json"] = &BaseResp{Result: true, Reason: content}
}

func (this *DrugController) fail(content interface{}) {
	this.Data["json"] = &BaseResp{Result: false, Reason: content}
}

func (this *DrugController) notFound() {
	this.Ctx.Output.SetStatus(404)
	this.Data["json"] = &BaseResp{Result: false, Reason: "Not Found"}
}
