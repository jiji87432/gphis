package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type T_Patient struct {
	P_ID       string  `json:"p_id" orm:"column(p_id)"`
	P_Name     string  `json:"p_name" orm:"column(p_name)"`
	P_Sex      string  `json:"p_sex" orm:"column(p_sex)"`
	P_Age      int     `json:"p_age" orm:"column(p_age)"`
	P_Birthday string  `json:"p_birthday" orm:"column(p_birthday)"`
	P_IDCard   string  `json:"p_idcard" orm:"column(p_idcard)"`
	P_Address  string  `json:"p_address" orm:"column(p_address)"`
	P_Contact  string  `json:"p_contact" orm:"column(p_contact)"`
	P_Height   float64 `json:"p_height" orm:"column(p_height)"`
	P_Weight   float64 `json:"p_weight" orm:"column(p_weight)"`
	P_GMS      string  `json:"p_gms" orm:"column(p_gms)"`
	P_SSS      string  `json:"p_sss" orm:"column(p_sss)"`
	P_MXBS     string  `json:"p_mxbs" orm:"column(p_mxbs)"`
}

func (this *T_Patient) Add() error {
	o := orm.NewOrm()
	rst, err := o.Raw("insert into patient (p_id,p_name,p_sex,p_age,p_birthday,p_idcard,p_address,p_contact,p_height,p_weight,p_gms,p_sss,p_mxbs) values (?,?,?,?,?,?,?,?,?,?,?,?,?)", this.P_ID, this.P_Name, this.P_Sex, this.P_Age, this.P_Birthday, this.P_IDCard, this.P_Address, this.P_Contact, this.P_Height, this.P_Weight, this.P_GMS, this.P_SSS, this.P_MXBS).Exec()
	if err != nil {
		return err
	}

	if rows, _ := rst.RowsAffected(); rows != 1 {
		return fmt.Errorf("受影响的行数为: %v", rows)
	}

	return nil
}

func (this *T_Patient) GetOneByID() (*T_Patient, error) {
	var ret *T_Patient
	o := orm.NewOrm()
	err := o.Raw("select * from patient where p_id='" + this.P_ID + "'").QueryRow(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (this *T_Patient) GetOneByName() ([]*T_Patient, error) {
	var ret []*T_Patient
	o := orm.NewOrm()
	_, err := o.Raw("select * from patient where p_name='" + this.P_Name + "'").QueryRows(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (this *T_Patient) DeleteByID() error {
	o := orm.NewOrm()
	rst, err := o.Raw("delete from patient where p_id='" + this.P_ID + "'").Exec()
	if err != nil {
		return err
	}

	if row, _ := rst.RowsAffected(); row != 1 {
		return fmt.Errorf("受影响行数为: %v", row)
	}

	return nil
}
