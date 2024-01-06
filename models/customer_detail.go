package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerDetail struct {
	gorm.Model
	LeadID                          uint64 `gorm:";NOT NULL"`
	Gender                          uint8  `gorm:"comment:1=>MALE, 2=>FEMALE"` // 1=>MALE, 2=>FEMALE
	Dob                             time.Time
	ResidenceAddress1               string
	ResidenceAddress2               string
	ResidenceAddress3               string
	ResidenceLandmark               string
	ResidenceStateID                uint16 `gorm:"comment:Residence state id"`        // Residence state id
	ResidenceCityID                 uint32 `gorm:"comment:Residence city id"`         // Residence city id
	ResidencePincode                uint32 `gorm:"comment:Current residence pincode"` // Current residence pincode
	ResidenceType                   int8   `gorm:"comment:master_residence_type"`     // master_residence_type
	ResidenceSince                  time.Time
	ResidenceWithFamily             uint8 `gorm:"default:0;comment:1=>YES"` // 1=>YES
	ResidenceEaadhaarAddressSame    uint8 `gorm:"default:0;comment:1=>YES"` // 1=>YES
	AadharSameAsResidenceAddress    string
	AadharNo                        string
	AadhaarAddress1                 string
	AadhaarAddress2                 string
	AadhaarAddress3                 string
	AadhaarLandmark                 string
	AadhaarStateID                  uint16
	AadhaarCityID                   uint32
	AadhaarPincode                  uint32
	CustomerReligionID              uint8
	FatherName                      string
	MotherName                      string
	CustomerDocsAvailable           uint8 `gorm:"default:0;comment:1=>Customer, 2=>Executive"` // 1=>Customer, 2=>Executive
	CustomerMaritalStatusID         uint8 `gorm:"comment:master_marital_status"`               // master_marital_status
	CustomerSpouseName              string
	CustomerSpouseOccupationID      uint8     `gorm:"comment:master_occupation"`        // master_occupation
	CustomerQualificationID         uint8     `gorm:"comment:master_qualification"`     // master_qualification
	CustomerDocumentUploadDatetime  time.Time `gorm:"comment:customer doc upload time"` // customer doc upload time
	CustomerExecutiveReloanFlag     uint8     `gorm:"comment:1=>YES, 2=>NO"`            // 1=>YES, 2=>NO
	CustomerExecutiveReloanRemark   string
	CustomerExecutiveReloanUserID   uint32
	CustomerExecutiveReloanDatetime time.Time
	CustomerActive                  uint8 `gorm:";default:1;NOT NULL"`
	CustomerDeleted                 uint8 `gorm:";default:0;NOT NULL"`
}
