package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gphis/models"
)

type PatientController struct {
	beego.Controller
}

func (this *PatientController) Post() {
	switch this.Ctx.Input.Param(":type") {
	case "add":
		this.PatientAdd()
	default:
		this.notFound()
	}

	this.ServeJSON()
}

func (this *PatientController) PatientAdd() {
	p := &model.T_Patient{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, p)
	if err != nil {
		this.fail(err.Error())
		return
	}

	rst, err := p.Add()
	if err != nil {
		this.fail(err.Error())
		return
	}

	this.succ("添加成功")
}

func (this *PatientController) succ(reason interface{}) {
	this.Data["json"] = &BaseResp{
		Result: true,
		Reason: reason,
	}
}

func (this *PatientController) fail(reason interface{}) {
	this.Data["json"] = &BaseResp{
		Result: false,
		Reason: reason,
	}
}
