package models

import (
	"github.com/astaxie/beego/orm"
)

type T_Drug struct {
	Drug_ID       string  `json:"drug_id" orm:"column(drug_id)"`
	Drug_Name     string  `json:"drug_name" orm:"column(drug_name)"`
	Drug_Pinyin   string  `json:"drug_pinyin" orm:"column(drug_pinyin)"`
	Drug_Barcode  string  `json:"drug_barcode" orm:"column(drug_barcode)"`
	Drug_Unit     string  `json:"drug_unit" orm:"column(drug_unit)"`
	Drug_U_Price  float64 `json:"drug_u_price" orm:"column(drug_u_price)"`
	Drug_Spec     string  `json:"drug_spec" orm:"column(drug_spec)"`
	Drug_Producer string  `json:"drug_producer" orm:"column(drug_producer)"`
	Drug_Mount    float64 `json:"drug_mount" orm:"column(drug_mount)"`
}

func (this *T_Drug) GetAll() ([]*T_Drug, error) {
	var ret []*T_Drug
	o := orm.NewOrm()
	_, err := o.Raw("select * from drug").QueryRows(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
