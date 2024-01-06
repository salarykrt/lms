package masters

import "gorm.io/gorm"

type MasterCompanyAccounts struct {
	gorm.Model
	DisbBankName          string `gorm:"NOT NULL"`
	DisbBankAccountNo     string `gorm:"NOT NULL"`
	DisbBankPaymentTypeID uint8  `gorm:"default:1;NOT NULL;comment:1=>only IMPS,2=>Only NEFT,3=>Both"` // 1=>only IMPS,2=>Only NEFT,3=>Both
	DisbBankImpsApiActive uint8  `gorm:"default:0;NOT NULL;comment:1=>Availabe"`                       // 1=>Availabe
	DisbBankNeftApiActive uint8  `gorm:"default:0;NOT NULL;comment:1=>Availabe"`                       // 1=>Availabe
	DisbBankActive        uint8  `gorm:"default:1;NOT NULL"`
	DisbBankDeleted       uint8  `gorm:"default:0;NOT NULL"`
}
