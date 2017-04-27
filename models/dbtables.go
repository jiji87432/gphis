package models

type T_Patient struct {
	P_ID       string `json:"p_id" orm:"column(p_id)"`
	P_Name     string `json:"p_name" orm:"column(p_name)"`
	P_Sex      string `json:"p_sex" orm:"column(p_sex)"`
	P_Age      int    `json:"p_age" orm:"column(p_age)"`
	P_Birthday string `json:"p_birthday" orm:"column(p_birthday)"`
	P_IDCard   string `json:"p_idcard" orm:"column(p_idcard)"`
	P_Address  string `json:"p_address" orm:"column(p_address)"`
	P_Contact  string `json:"p_contact" orm:"column(p_contact)"`
	P_Height   int    `json:"p_height" orm:"column(p_height)"`
	P_Weight   int    `json:"p_weight" orm:"column(p_weight)"`
	P_GMS      string `json:"p_gms" orm:"column(p_gms)"`
	P_SSS      string `json:"p_sss" orm:"column(p_sss)"`
	P_MXBS     string `json:"p_mxbs" orm:"column(p_mxbs)"`
}

type T_Operator struct {
	O_ID     string `json:"o_id" orm:"column(o_id)"`
	O_Name   string `json:"o_name" orm:"column(o_name)"`
	O_Pinyin string `json:"o_pinyin" orm:"column(o_pinyin)"`
	O_Permit string `json:"o_permit" orm:"column(o_permit)"`
	O_Pwd    string `json:"o_pwd" orm:"column(o_pwd)"`
}

type T_Drug_Item struct {
	Drug_Item_ID string  `json:"drug_item_id" orm:"column(drug_item_id)"`
	Drug_ID      string  `json:"drug_id" orm:"column(drug_id)"`
	Mount        float64 `json:"mount" orm:"column(mount)"`
}

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

type T_Bills struct {
	B_ID          string  `json:"b_id" orm:"column(b_id)"`
	Diag_ID       string  `json:"diag_id" orm:"column(diag_id)"`
	Total_Expense float64 `json:"total_expense" orm:"column(total_expense)"`
	O_ID          string  `json:"o_id" orm:"column(o_id)"`
	B_Date        string  `json:"b_date" orm:"column(b_date)"`
}
