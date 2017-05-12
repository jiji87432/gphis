package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gphis/models"
)

type OperatorController struct {
	beego.Controller
}

func (this *OperatorController) Post() {
	switch this.Ctx.Input.Param(":type") {
	case "get_one_by_pinyin":
		this.GetOneByPinyin()
	default:
		this.notFound()
	}

	this.ServeJSON()
}

func (this *OperatorController) GetOneByPinyin() {
	req := &models.T_Operator{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, req)
	if err != nil {
		this.fail(err.Error())
		return
	}

	if req.O_Pinyin == "" {
		this.fail("未提供拼音简码")
		return
	}

	q, err := req.GetOneByPinyin()
	if err != nil {
		this.fail(err.Error())
		return
	}

	this.succ(q)
}

func (this *OperatorController) succ(reason interface{}) {
	this.Data["json"] = &BaseResp{Result: true, Reason: reason}
}

func (this *OperatorController) fail(reason interface{}) {
	this.Data["json"] = &BaseResp{Result: false, Reason: reason}
}

func (this *OperatorController) notFound() {
	this.Ctx.Output.SetStatus(404)
	this.Data["json"] = &BaseResp{Result: false, Reason: "Not Found"}
}
