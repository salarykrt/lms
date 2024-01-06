package models

import (
	"time"

	"gorm.io/gorm"
)

type CifCustomer struct {
	gorm.Model
	CifNumber                      string
	CifFirstName                   string
	CifMiddleName                  string
	CifSurName                     string
	CifGender                      uint8 `gorm:"comment:1=>Male,2=>Female"` // 1=>Male,2=>Female
	CifDob                         time.Time
	CifPancard                     string
	CifPancardVerified             uint8 `gorm:"default:0"`
	CifPancardVerifiedOn           time.Time
	CifPersonalEmail               string
	CifOfficeEmail                 string
	CifMobile                      uint64
	CifAlternateMobile             uint64
	CifResidenceAddress1           string
	CifResidenceAddress2           string
	CifResidenceAddress3           string
	CifResidenceLandmark           string
	CifResidenceCityID             uint32 `gorm:"comment:master_city"`    // master_city
	CifResidenceStateID            uint16 `gorm:"comment:master_state"`   // master_state
	CifResidencePincode            uint32 `gorm:"comment:master_pincode"` // master_pincode
	CifResidenceSince              time.Time
	CifResidenceTypeID             uint8 `gorm:"comment:master_residence_type"` // master_residence_type
	CifResidenceResidingWithFamily uint8 `gorm:"comment:1=>YES, 2=>NO"`         // 1=>YES, 2=>NO
	CifAadhaarNo                   uint64
	CifAadhaarRefKey               string
	CifOfficeAddress1              string
	CifOfficeAddress2              string
	CifOfficeAddress3              string
	CifOfficeAddressLandmark       string
	CifOfficeCityID                uint32 `gorm:"comment:master_city"`    // master_city
	CifOfficeStateID               uint16 `gorm:"comment:master_state"`   // master_state
	CifOfficePincode               uint32 `gorm:"comment:master_pincode"` // master_pincode
	CifOfficeWorkingSince          time.Time
	CifOfficeDesignation           string
	CifOfficeDepartment            string
	CifCompanyName                 string
	CifCompanyWebsite              string
	CifCompanyTypeID               uint8 `gorm:"comment:master_company_type"` // master_company_type
	CifAadhaarSameAsResidence      uint8 `gorm:";default:0;comment:1=>YES"`   // 1=>YES
	CifAadhaarAddress1             string
	CifAadhaarAddress2             string
	CifAadhaarAddress3             string
	CifAadhaarLandmark             string
	CifAadhaarCityID               uint32 `gorm:"comment:master_city"`                  // master_city
	CifAadhaarStateID              uint16 `gorm:"comment:master_state"`                 // master_state
	CifAadhaarPincode              uint32 `gorm:"comment:master_pincode"`               // master_pincode
	CifLoanIsDisbursed             uint8  `gorm:";default:0;comment:1=>Loan dsibursed"` // 1=>Loan dsibursed
	CifLoanDisbursedCount          uint32 `gorm:";default:0;NOT NULL"`
	CifCreatedBy                   uint32
	CifUpdatedBy                   uint32
	CifIncomeTypeID                uint8 `gorm:"comment:1=>salaried;2=>self-employed"` // 1=>salaried;2=>self-employed
	CifDigitalEkycFlag             uint8 `gorm:";default:0;comment:1=>Yes"`            // 1=>Yes
	CifDigitalEkycDatetime         time.Time
	CifActive                      uint8 `gorm:";default:1;NOT NULL"`
	CifDeleted                     uint8 `gorm:";default:0;NOT NULL"`
	CifFatherName                  string
	CifReligionID                  uint8 `gorm:"comment:comment:1=>Yes"`        // master_religion
	CifMaritalStatusID             uint8 `gorm:"comment:master_marital_status"` // master_marital_status
	CifSpouseName                  string
	CifSpouseOccupationID          uint8 `gorm:"comment:master_occupation"` // master_occupation
	CifQualificationID             uint8 `gorm:"comment:master_occupation"` // 	master_occupation
	CifOccupationID                uint8 `gorm:"comment:master_occupation"` // master_occupation
	CifExecutiveReloanFlag         uint8 `gorm:"comment:1=>YES, 2=>NO"`     // 1=>YES, 2=>NO
	CifExecutiveReloanRemark       string
	CifExecutiveReloanUserID       uint32
	CifExecutiveReloanDatetime     time.Time
}

func (m *CifCustomer) TableName() string {
	return "cif_customer"
}
