package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gphis/utils"
	"time"
)

type T_Drug_Item struct {
	Drug_Item_ID string  `json:"drug_item_id" orm:"column(drug_item_id)"`
	Drug_ID      string  `json:"drug_id" orm:"column(drug_id)"`
	Mount        float64 `json:"mount" orm:"column(mount)"`
}

type DrugItemDetail struct {
	Drug_ID       string  `json:"drug_id" orm:"column(drug_id)"`
	Drug_Item_ID  string  `json:"drug_item_id" orm:"column(drug_item_id)"`
	Mount         float64 `json:"mount" orm:"column(mount)"`
	Drug_Name     string  `json:"drug_name" orm:"column(drug_name)"`
	Drug_Pinyin   string  `json:"drug_pinyin" orm:"column(drug_pinyin)"`
	Drug_Barcode  string  `json:"drug_barcode" orm:"column(drug_barcode)"`
	Drug_Unit     string  `json:"drug_unit" orm:"column(drug_unit)"`
	Drug_U_Price  float64 `json:"drug_u_price" orm:"column(drug_u_price)"`
	Drug_Spec     string  `json:"drug_spec" orm:"column(drug_spec)"`
	Drug_Producer string  `json:"drug_producer" orm:"column(drug_producer)"`
	Drug_Mount    float64 `json:"drug_mount" orm:"column(drug_mount)"`
}

func (this *T_Diagnose) GetDrugItemDetail() ([]*DrugItemDetail, error) {
	var ret []*DrugItemDetail
	o := orm.NewOrm()
	// 定义新的ormer, 使用Raw函数执行原生SQL语句, 进行多表联合查询
	_, err := o.Raw("select drug.drug_id,drug.drug_name,drug.drug_pinyin,drug.drug_barcode,drug.drug_unit,drug.drug_u_price,drug.drug_spec,drug.drug_producer,drug.drug_mount,drug_item.drug_item_id,drug_item.mount from drug,drug_item where drug.drug_id=drug_item.drug_id and drug_item.drug_item_id='" + this.Drug_Item_ID + "'").QueryRows(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type T_Dispense_Item struct {
	Dis_Item_ID string  `json:"dis_item_id" orm:"column(dis_item_id)"`
	Dis_ID      string  `json:"dis_id" orm:"column(dis_id)"`
	Mount       float64 `json:"mount" orm:"column(mount)"`
}

type T_Dispense struct {
	Dis_ID      string  `json:"dis_id" orm:"column(dis_id)"`
	Dis_Name    string  `json:"dis_name" orm:"column(dis_name)"`
	Dis_Pinyin  string  `json:"dis_pinyin" orm:"column(dis_pinyin)"`
	Dis_Unit    string  `json:"dis_unit" orm:"column(dis_unit)"`
	Dis_U_Price float64 `json:"dis_u_price" orm:"column(dis_u_price)"`
}

type DispenseItemDetail struct {
	Dis_ID      string  `json:"dis_id" orm:"column(dis_id)"`
	Dis_Name    string  `json:"dis_name" orm:"column(dis_name)"`
	Dis_Pinyin  string  `json:"dis_pinyin" orm:"column(dis_pinyin)"`
	Dis_Unit    string  `json:"dis_unit" orm:"column(dis_unit)"`
	Dis_U_Price float64 `json:"dis_u_price" orm:"column(dis_u_price)"`
	Dis_Item_ID string  `json:"dis_item_id" orm:"column(dis_item_id)"`
	Mount       float64 `json:"mount" orm:"column(mount)"`
}

func (this *T_Diagnose) GetDispenseItemDetail() ([]*DispenseItemDetail, error) {
	var ret []*DispenseItemDetail
	o := orm.NewOrm()
	_, err := o.Raw("select dispense.dis_id,dispense.dis_name,dispense.dis_pinyin,dispense.dis_unit,dispense.dis_u_price,dispense_item.dis_item_id,dispense_item.mount from dispense,dispense_item where dispense.dis_id=dispense_item.dis_id and dispense_item.dis_item_id='" + this.Dis_Item_ID + "'").QueryRows(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type T_Diagnose struct {
	Diag_ID       string `json:"diag_id" orm:"column(diag_id)"`
	Diag_Date     string `json:"diag_date" orm:"column(diag_date)"`
	P_ID          string `json:"p_id" orm:"column(p_id)"`
	O_ID          string `json:"o_id" orm:"column(o_id)"`
	Drug_Item_ID  string `json:"drug_item_id" orm:"column(drug_item_id)"`
	Dis_Item_ID   string `json:"dis_item_id" orm:"column(dis_item_id)"`
	Diag_Complain string `json:"diag_complain" orm:"column(diag_complain)"`
	Diag_Checks   string `json:"diag_checks" orm:"column(diag_checks)"`
	Diag_Diagnose string `json:"diag_diagnose" orm:"column(diag_diagnose)"`
}

func (this *T_Diagnose) QueryOne() (*T_Diagnose, error) {
	ret := &T_Diagnose{}
	o := orm.NewOrm()
	err := o.Raw("select * from diagnose where diagnose.diag_id='" + this.Diag_ID + "'").QueryRow(ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type T_Bills struct {
	B_ID          string  `json:"b_id" orm:"column(b_id)"`
	Diag_ID       string  `json:"diag_id" orm:"column(diag_id)"`
	Total_Expense float64 `json:"total_expense" orm:"column(total_expense)"`
	O_ID          string  `json:"o_id" orm:"column(o_id)"`
	B_Date        string  `json:"b_date" orm:"column(b_date)"`
}

func (this *T_Diagnose) SquareAccount(oid string) (*T_Bills, error) {
	if this.Diag_ID == "" {
		return nil, fmt.Errorf("接诊单ID为空")
	}

	// 查询指定接诊单是否存在
	diag, err := this.QueryOne()
	if err != nil {
		return nil, err
	}

	if diag.Diag_ID == "" {
		return nil, fmt.Errorf("不存在指定接诊单")
	}

	// 查询指定接诊单的处方信息
	drugItem, err := this.GetDrugItemDetail()
	if err != nil {
		return nil, err
	}

	// 查询指定接诊单的处置信息
	disItem, err := this.GetDispenseItemDetail()
	if err != nil {
		return nil, err
	}

	// 判定是否开具过收费项目
	if len(drugItem) == 0 && len(disItem) == 0 {
		return nil, fmt.Errorf("尚未开具收费项目")
	}

	// 声明总费用变量
	var totalExpense float64

	// 统计处方总费用
	if len(drugItem) != 0 {
		for _, v := range drugItem {
			totalExpense += v.Drug_U_Price * v.Mount
		}
	}

	// 统计处置总费用
	if len(disItem) != 0 {
		for _, v := range disItem {
			totalExpense += v.Dis_U_Price * v.Mount
		}
	}

	bill := &T_Bills{
		B_ID:          utils.GetSerial(),
		Diag_ID:       this.Diag_ID,
		Total_Expense: totalExpense,
		O_ID:          oid,
		B_Date:        time.Now().String(),
	}

	// 执行原生SQL INSERT语句
	o := orm.NewOrm()
	rst, err := o.Raw("insert into bills (b_id,diag_id,total_expense,o_id,b_date) values (?,?,?,?,?)", bill.B_ID, bill.Diag_ID, bill.Total_Expense, bill.O_ID, bill.B_Date).Exec()
	if err != nil {
		return nil, err
	}

	// 插入操作后受影响的行数若不为1, 则报错
	if rows, _ := rst.RowsAffected(); rows != 1 {
		return nil, fmt.Errorf("RowsAffected的值为%v", rows)
	}

	// 插入操作正确, 结算流程结束
	return bill, nil
}
