package models

import "gorm.io/gorm"

type CustomerBanking struct {
	gorm.Model
	UserID                 uint32
	LeadID                 uint64
	BankName               string `gorm:"NOT NULL"`
	IfscCode               string `gorm:"NOT NULL"`
	Branch                 string
	BeneficiaryName        string
	Account                string `gorm:"NOT NULL"`
	ConfirmAccount         string `gorm:"NOT NULL"`
	AccountTypeID          uint8  `gorm:"comment:master_company_accounts"`    // master_company_accounts
	AccountStatusID        uint16 `gorm:"comment:master_bank_account_status"` // master_bank_account_status id
	Remark                 string
	CreatedBy              uint32
	UpdatedBy              uint32
	CustomerBankingActive  uint8 `gorm:"default:1;NOT NULL"`
	CustomerBankingDeleted uint8 `gorm:"default:0;NOT NULL"`
}

func (m *CustomerBanking) TableName() string {
	return "customer_banking"
}
