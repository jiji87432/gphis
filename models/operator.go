package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type T_Operator struct {
	O_ID     string `json:"o_id" orm:"column(o_id)"`
	O_Name   string `json:"o_name" orm:"column(o_name)"`
	O_Pinyin string `json:"o_pinyin" orm:"column(o_pinyin)"`
	O_Permit string `json:"o_permit" orm:"column(o_permit)"`
	O_Pwd    string `json:"o_pwd" orm:"column(o_pwd)"`
}

func (this *T_Operator) GetOneByPinyin() (*T_Operator, error) {
	var ret *T_Operator
	o := orm.NewOrm()
	err := o.Raw("select * from operator where o_pinyin = '" + this.O_Pinyin + "'").QueryRow(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (this *T_Operator) GetOneByID() (*T_Operator, error) {
	var ret *T_Operator
	o := orm.NewOrm()
	err := o.Raw("select * from operator where o_id='" + this.O_ID + "'").QueryRow(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (this *T_Operator) Add() error {
	o := orm.NewOrm()
	rst, err := o.Raw("insert into operator (o_id,o_name,o_pinyin,o_permit,o_pwd) values (?,?,?,?,?)", this.O_ID, this.name, this.O_Pinyin, this.O_Permit, this.O_Pwd).Exec()
	if err != nil {
		return err
	}

	if row, _ := rst.RowsAffected(); row != 1 {
		return fmt.Errorf("影响行数为: %v", row)
	}

	return nil
}

func (this *T_Operator) UpdateByID() error {
	o := orm.NewOrm()
	rst, err := o.Raw("update operator set o_name=?,o_pinyin=?,o_permit=?,o_pwd=? where o_id=?", this.O_Name, this.O_Pinyin, this.O_Permit, this.O_Pwd, this.O_ID).Exec()
	if err != nil {
		return err
	}

	if row, _ := rst.RowsAffected(); row != 1 {
		return fmt.Errorf("受影响行数为: %v", row)
	}

	return nil
}

func (this *T_Operator) DeleteByID() error {
	o := orm.NewOrm()
	rst, err := o.Raw("delete from operator where o_id=?", this.O_ID).Exec()
	if err != nil {
		return err
	}

	if row, _ := rst.RowsAffected(); row != 1 {
		return fmt.Errorf("受影响行数为: %v", row)
	}

	return nil
}
