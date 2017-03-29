package models

type T_DrugInfo struct {
	DrugId       string  `json:"drugID" orm:"column(drugID)"`
	ProducerId   string  `json:"producerID" orm:"column(producerID)"`
	DrugName     string  `json:"drugName" orm:"column(drugName)"`
	DrugMnemonic string  `json:"drugMnemonic" orm:"column(drugMnemonic)"`
	Barcode      string  `json:"barcode" orm:"column(barcode)"`
	Unit         string  `json:"unit" orm:"column(unit)"`
	ProSize      string  `json:"proSize" orm:"column(proSize)"`
	PriceIn      float64 `json:"priceIn" orm:"column(priceIn)"`
	PriceOut     float64 `json:"priceOut" orm:"column(priceOut)"`
}

type D_IDCard struct {
	IDNum string `json:"idNum" orm:"column(idNum)"`
	Place string `json:"place" orm:"column(place)"`
}

type T_Acceptance struct {
	AcceptanceID   string `json:"acceptanceID" orm:"column(acceptanceID)"`
	PatientID      string `json:"patientID" orm:"column(patientID)"`
	OperatorID     string `json:"operatorID" orm:"column(operatorID)"`
	ChiefComplain  string `json:"chiefComplain" orm:"column(chiefComplain)"`
	MedicalHistory string `json:"medicalHistory" orm:"column(medicalHistory)"`
	Checks         string `json:"checks" orm:"column(checks)"`
	Treats         string `json:"treats" orm:"column(treats)"`
	Diagnose       string `json:"diagnose" orm:"column(diagnose)"`
}

type T_Balance struct {
	BalanceID    string  `json:"balanceID" orm:"column(balanceID)"`
	TypeID       string  `json:"typeID" orm:"column(typeID)"`
	TotalMount   float64 `json:"totalMount" orm:"column(totalMount)"`
	OperatorID   string  `json:"operatorID" orm:"column(operatorID)"`
	AcceptanceID string  `json:"acceptanceID" orm:"column(acceptanceID)"`
}

type T_ChargeType struct {
	TypeID   string `json:"typeID" orm:"column(typeID)"`
	TypeName string `json:"typeName" orm:"column(typeName)"`
}

type T_DetailList struct {
	AcceptanceID string  `json:"acceptanceID" orm:"column(acceptanceID)"`
	ItemID       string  `json:"itemID" orm:"column(itemID)"`
	ItemName     string  `json:"itemName" orm:"column(itemName)"`
	Unit         string  `json:"unit" orm:"column(unit)"`
	UnitPrice    float64 `json:"unitPrice" orm:"column(unitPrice)"`
	Mount        float64 `json:"mount" orm:"column(mount)"`
	TotalPrice   float64 `json:"totalPrice" orm:"column(totalPrice)"`
	OperatorID   string  `json:"operatorID" orm:"column(operatorID)"`
	TypeID       string  `json:"typeID" orm:"column(typeID)"`
}

type T_Operator struct {
	OperatorID   string `json:"operatorID" orm:"column(operatorID)"`
	Name         string `json:"name" orm:"column(name)"`
	MnemonicCode string `json:"mnemonicCode" orm:"column(mnemonicCode)"`
	Password     string `json:"password" orm:"column(password)"`
}

type T_Patient struct {
	PatientID string `json:"patientID" orm:"column(patientID)"`
	Name      string `json:"name" orm:"column(name)"`
	Sex       string `json:"sex" orm:"column(sex)"`
	Age       int    `json:"age" orm:"column(age)"`
	Birthday  string `json:"birthday" orm:"column(birthday)"`
	IDNum     string `json:"idnum" orm:"column(idnum)"`
	Address   string `json:"address" orm:"column(address)"`
	Phone     string `json:"phone" orm:"column(phone)"`
}

type T_Producer struct {
	ProducerID   string `json:"producerID" orm:"column(producerID)"`
	ProducerName string `json:"producerName" orm:"column(producerName)"`
	Contact      string `json:"contact" orm:"column(contact)"`
}

type T_StockInfo struct {
	StockID    string  `json:"stockID" orm:"column(stockID)"`
	DrugID     string  `json:"drugID" orm:"column(drugID)"`
	UnitPrice  float64 `json:"unitPrice" orm:"column(unitPrice)"`
	MountIn    float64 `json:"mountIn" orm:"column(mountIn)"`
	Stock      float64 `json:"stock" orm:"column(stock)"`
	OperatorID string  `json:"operatorID" orm:"column(operatorID)"`
	Date       string  `json:"date" orm:"column(date)"`
}

type T_Treatment struct {
	TreatmentID   string  `json:"treatmentID" orm:"column(treatmentID)"`
	TreatmentName string  `json:"treatmentName" orm:"column(treatmentName)"`
	Unit          string  `json:"unit" orm:"column(unit)"`
	UnitPrice     float64 `json:"unitPrice" orm:"column(unitPrice)"`
}
