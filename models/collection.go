package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	LeadID                              uint64          `gorm:"NOT NULL"`
	LoanNo                              string          `gorm:"NOT NULL"`
	PaymentModeID                       uint8           `gorm:"comment:master_payment_mode"` // master_payment_mode
	ReceivedAmount                      decimal.Decimal `gorm:"NOT NULL;type:decimal(10,2)"`
	RefrenceNo                          string
	RepaymentStatusId                   uint16 `gorm:"comment:master_status"`           // master_status
	CompanyAccountID                    uint32 `gorm:"comment:master_company_accounts"` // master_company_accounts
	DocFileName                         string
	Discount                            decimal.Decimal `gorm:"NOT NULL;type:decimal(10,2)"`
	Refund                              decimal.Decimal `gorm:"type:decimal(10,2)"`
	DateOfRecived                       time.Time
	PaymentVerification                 uint8           `gorm:"default:0;NOT NULL;comment:0=Pending, 1=Approved, 2=Reject"` // 0=Pending, 1=Approved, 2=Reject
	Sattelment                          decimal.Decimal `gorm:"type:decimal(10,2);default:0;NOT NULL"`
	Remarks                             string          `gorm:"NOT NULL"`
	Noc                                 string          `gorm:"default:0"`
	IP                                  string          `gorm:"NOT NULL"`
	CollectionExecutivePaymentCreatedOn time.Time       `gorm:"NOT NULL"`
	CollectionExecutiveUserID           uint32
	ClosureUserID                       uint
	ClosurePaymentUpdatedOn             time.Time
	ClosureRemarks                      string
	CollectionType                      string `gorm:"comment:0=collection default,1=collection new,2=collection reschedule"` // 0=collection default,1=collection new,2=collection reschedule
	CollectionActive                    uint   `gorm:"default:1;NOT NULL"`
	CollectionDeleted                   uint   `gorm:"default:0;NOT NULL"`
}

func (m *Collection) TableName() string {
	return "collection"
}
