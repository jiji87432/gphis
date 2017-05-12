package controllers

import (
	"github.com/astaxie/beego"
)

type PatientController struct {
	beego.Controller
}

func (this *PatientController) succ(reason interface{}) {
	this.Data["json"] = &BaseResp{Result: true, Reason: reason}
}

func (this *PatientController) fail(reason interface{}) {
	this.Data["json"] = &BaseResp{Result: false, Reason: reason}
}

func (this *PatientController) notFound() {
	this.Ctx.Output.SetStatus(404)
	this.Data["json"] = &BaseResp{Result: false, Reason: "Not Found"}
}
