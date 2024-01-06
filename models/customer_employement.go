package models

import (
	"time"
)

type CustomerEmployment struct {
	ID                      uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	LeadID                  uint64 `gorm:"NOT NULL"`
	EmpName                 string
	EmpAddress1             string
	EmpAddress2             string
	EmpAddress3             string
	EmpAddressLandmark      string
	EmpCityID               uint32
	EmpStateID              uint16
	EmpPincode              uint32
	EmpResidenceSince       time.Time
	EmpDesignation          string
	EmpDepartment           string
	EmpSalaryModeID         uint8 `gorm:"comment:master_salary_mode"`           // master_salary_mode
	EmpIncomeType           uint8 `gorm:"comment:1=>salaried;2=>self-employed"` // 1=>salaried;2=>self-employed
	EmpEmployerTypeID       uint8 `gorm:"comment:master_company_type"`          // master_company_type
	EmpPresentServiceTenure uint8
	EmpMonthlyIncome        float64
	EmpIndustryID           uint8 `gorm:"comment:master_industry"` // master_industry
	EmpStatusID             uint8 `gorm:"comment:1=>YES, 2=>NO"`   // 1=>YES, 2=>NO
	EmpCreatedBy            uint32
	EmpUpdatedBy            uint32
	EmpEmail                string
	EmpOccupationID         uint8 `gorm:"comment:master_occupation"`         // master_occupation
	EmpWorkModeID           uint8 `gorm:"comment:1=>Office,2=>Home,3=>Both"` // 1=>Office,2=>Home,3=>Both
	EmpActive               uint8 `gorm:"default:1;NOT NULL"`
	EmpDeleted              uint8 `gorm:"default:0;NOT NULL"`
}

func (m *CustomerEmployment) TableName() string {
	return "customer_employment"
}
