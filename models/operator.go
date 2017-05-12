package models

import (
	"github.com/astaxie/beego/orm"
)

type T_Operator struct {
	O_ID     string `json:"o_id" orm:"column(o_id)"`
	O_Name   string `json:"o_name" orm:"column(o_name)"`
	O_Pinyin string `json:"o_pinyin" orm:"column(o_pinyin)"`
	O_Permit string `json:"o_permit" orm:"column(o_permit)"`
	O_Pwd    string `json:"o_pwd" orm:"column(o_pwd)"`
}

func (this *T_Operator) GetOne() (*T_Operator, error) {
	var ret *T_Operator
	o := orm.NewOrm()
	err := o.Raw("select * from operator where o_pinyin = '" + this.O_Pinyin + "'").QueryRow(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
