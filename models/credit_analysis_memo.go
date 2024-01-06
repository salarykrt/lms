package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CreditAnalysisMemo struct {
	gorm.Model
	LeadID                         uint64 `gorm:"NOT NULL"`
	Ntc                            uint8  `gorm:"comment:1=>YES, 2=>NO"` // 1=>YES, 2=>NO
	RunOtherPdLoan                 uint8  `gorm:"comment:1=>YES, 2=>NO"` // 1=>YES, 2=>NO
	DelayOtherLoan30Days           uint8  `gorm:"comment:1=>YES, 2=>NO"` // 1=>YES, 2=>NO
	JobStability                   uint8  `gorm:"comment:1=>YES, 2=>NO"` // 1=>LOW, 2=>MEDIUM, 3=>HIGH
	CityCategory                   string
	SalaryCredit1Date              time.Time
	SalaryCredit1Amount            uint32
	SalaryCredit2Date              time.Time
	SalaryCredit2Amount            uint32
	SalaryCredit3Date              time.Time
	SalaryCredit3Amount            uint32
	NextPayDate                    time.Time
	MedianSalaryAmount             uint32
	SalaryVariance                 string
	BorrowerAge                    string
	EligibleFoirPercentage         decimal.Decimal `gorm:"type:decimal(10,2)"`
	EligibleLoan                   decimal.Decimal `gorm:"type:decimal(10,2)"`
	FinalFoirPercentage            decimal.Decimal `gorm:"type:decimal(10,2)"`
	FoirEnhancedBy                 string
	ProcessingFeePercent           decimal.Decimal `gorm:"type:decimal(10,2)"`
	Roi                            decimal.Decimal `gorm:"type:decimal(10,2)"`
	LoanRecommended                uint32
	DisbursalDate                  time.Time
	RepaymentDate                  time.Time
	AdminFee                       decimal.Decimal `gorm:"comment:without gst included;type:decimal(10,2)"` // without gst included
	AdminFeeGst                    decimal.Decimal `gorm:"comment:calculated GST;type:decimal(10,2)"`       // calculated GST
	TotalAdminFee                  decimal.Decimal `gorm:"comment:net pf with gst;type:decimal(10,2)"`      // net pf with gst
	Tenure                         uint16
	NetDisbursalAmount             decimal.Decimal `gorm:"type:decimal(10,2)"`
	RepaymentAmount                decimal.Decimal `gorm:"type:decimal(10,2)"`
	PanelRoi                       decimal.Decimal `gorm:"type:decimal(10,2)"`
	B2BNumber                      uint16
	Remark                         string
	CreatedBy                      uint32
	UpdatedBy                      uint32
	CamSanctionLetterFileName      string
	CamSanctionLetterEsginTypeID   uint32 `gorm:"default:0;comment:1=>aadhaar eSign"` // 1=>aadhaar eSign
	CamSanctionLetterEsginFileName string
	CamSanctionLetterEsginOn       time.Time
	CamSanctionLetterIPAddress     string
	CamSanctionLetterEsginCount    uint16 `gorm:"default:0"`
	CamRiskProfileID               uint8  `gorm:"comment:1=>LOW, 2=>MEDIUM, 3=>HIGH"` // 1=>LOW, 2=>MEDIUM, 3=>HIGH
	CamRiskScore                   uint16
	CamAdvanceInterestAmount       decimal.Decimal `gorm:"default:0;type:decimal(10,2)"`
	CamAppraisedObligations        decimal.Decimal `gorm:"default:0;type:decimal(10,2)"`
	CamAppraisedMonthlyIncome      decimal.Decimal `gorm:"default:0;type:decimal(10,2)"`
	CamSanctionRemarks             string
	CamProcessingFeeGstTypeID      uint8 `gorm:"comment:1=>Inclusive,2=>Exclusive"` // 1=>Inclusive,2=>Exclusive
	CamActive                      uint8 `gorm:"default:1;NOT NULL"`
	CamDeleted                     uint8 `gorm:"default:0;NOT NULL"`
}

func (m *CreditAnalysisMemo) TableName() string {
	return "credit_analysis_memo"
}
