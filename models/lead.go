package models

import (
	LeadStage "salarykart/ctypes/leadstage"
	LeadStatus "salarykart/ctypes/leadstatus"
	"salarykart/ctypes/usertype"
	"time"

	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	CustomerID                                string
	CompanyID                                 uint8 `gorm:"default:1;NOT NULL"`
	ProductID                                 uint8 `gorm:"default:1;NOT NULL"`
	ApplicationNo                             string
	LoanNo                                    string
	EnduseID                                  uint8 `gorm:"comment:master_enduse"` // master_enduse
	UserType                                  usertype.UserType
	FirstName                                 string
	MiddleName                                string
	LastName                                  string
	Pancard                                   string
	Mobile                                    uint64
	Otp                                       int16
	MobileVerified                            uint8 `gorm:"default:0;NOT NULL;comment:1=>VERIFIED"` // 1=>VERIFIED
	AlternateMobile                           uint64
	Email                                     string
	AlternateEmail                            string
	AppliedAmount                             uint32
	AppliedTenure                             uint32
	CibilScore                                int32
	CheckCibilStatus                          int8 `gorm:"comment:1=>FETCHED"` // 1=>FETCHED
	Obligations                               uint32
	Promocode                                 string
	LeadBranchID                              uint32
	StateID                                   int16
	CityID                                    uint32
	Pincode                                   uint32
	TermAndCondition                          int8 `gorm:"default:0;comment:1=>AGREED"` // 1=>AGREED
	Coordinates                               string
	Status                                    LeadStatus.LeadStatus `gorm:"NOT NULL"`
	LeadStatusID                              uint16
	Stage                                     LeadStage.LeadStage `gorm:"NOT NULL"`
	Remark                                    string
	Source                                    string
	LeadDataSourceID                          int8 `gorm:"comment:master_data_source"` // master_data_source
	UtmSource                                 string
	UtmMedium                                 string
	UtmTerm                                   string
	UtmCampaign                               string
	IP                                        string
	LeadEntryDate                             time.Time
	LeadReferenceNo                           string
	ApplicationStatus                         uint8 `gorm:"default:0;comment:0=>Can Recommend, 1=>can't recommend"` // 0=> Can Recommend, 1=>can't recommend
	LeadScreenerAssignUserID                  uint32
	LeadScreenerAssignDatetime                time.Time
	LeadScreenerRecommendDatetime             time.Time
	LeadCreditAssignUserID                    uint32
	LeadCreditAssignDatetime                  time.Time
	LeadCreditRecommendDatetime               time.Time
	LeadCreditheadAssignUserID                uint32
	LeadCreditheadAssignDatetime              time.Time
	LeadCreditApproveUserID                   uint32
	LeadCreditApproveDatetime                 time.Time
	LeadScmAssignUserID                       uint32
	LeadScmAssignDatetime                     time.Time
	LeadCfeAssignUserID                       uint32
	LeadCfeAssignDatetime                     time.Time
	LeadClosureAssignUserID                   uint32
	LeadClosureAssignDatetime                 time.Time
	LeadDisbursalAssignUserID                 uint32
	LeadDisbursalAssignDatetime               time.Time
	LeadDisbursalRecommendDatetime            time.Time
	LeadDisbursalApproveUserID                uint32
	LeadDisbursalApproveDatetime              time.Time
	LeadFinalDisbursedDate                    time.Time
	LeadRejectedReasonID                      uint32 `gorm:"comment:master_rejection"` // master_rejection
	LeadRejectedUserID                        uint32
	LeadRejectedDatetime                      time.Time
	LeadBlackListFlag                         uint8 `gorm:"default:0;NOT NULL;comment:1=>YES"`                           // 1=>YES
	LeadStpFlag                               uint8 `gorm:"default:0;NOT NULL;comment:Lead Straight-Through Processing"` // Lead Straight-Through Processing
	LeadRejectedAssignUserID                  uint64
	LeadRejectedAssignDatetime                time.Time
	LeadRejectedAssignCounter                 uint16
	LeadJourneyTypeID                         uint8 `gorm:"comment:1=>Web,2=>App"`                       // 1=>Web,2=>App
	LeadJourneyStageID                        uint8 `gorm:"comment: master_journey_stage"`               // master_journey_stage
	LeadJourneyBreStatus                      uint8 `gorm:"comment:1=>Approved,2=>Referred,3=>Rejected"` // 1=>Approved,2=>Referred,3=>Rejected
	LeadLoanLastFollowupID                    uint64
	LeadLoanLastFollowupTypeID                uint8
	LeadLoanLastFollowupUserID                uint32
	LeadLoanLastFollowupDatetime              time.Time
	LeadPreCollectionExecutiveAssignUserID1   uint32
	LeadPreCollectionExecutiveAssignDatetime1 time.Time
	LeadPreCollectionExecutiveAssignUserID2   uint32
	LeadPreCollectionExecutiveAssignDatetime2 time.Time
	LeadCollectionExecutiveAssignUserID1      uint32
	LeadCollectionExecutiveAssignDatetime1    time.Time
	LeadCollectionExecutiveAssignUserID2      uint32
	LeadCollectionExecutiveAssignDatetime2    time.Time
	LeadActive                                uint8 `gorm:"default:1;NOT NULL"`
	LeadDeleted                               uint8 `gorm:"default:0;NOT NULL"`
}

// NewLeadFromSchema converts the schema.Lead to models.Lead
// func NewLeadFromSchema(lead schema.Lead) Lead {
// 	return Lead{
// 		LoanAmount:       decimal.NewFromInt(int64(lead.LoanAmount)),
// 		FirstName:        lead.Name,
// 		Pancard:          lead.PanCard,
// 		Obligations:      decimal.NewFromInt(int64(lead.Obligations)),
// 		Mobile:           ctypes.ConvertToUint(lead.Mobile),
// 		Email:            lead.EmailPersonal,
// 		AlternateEmail:   lead.EmailOffice,
// 		Source:           lead.Source,
// 		LeadDataSourceId: lead.SourceId,
// 		CityId:           lead.CityId,
// 		StateId:          lead.StateId,
// 		Pincode:          ctypes.ConvertToUint(lead.Pincode),
// 		UserType:         usertype.New,
// 		LeadEntryDate:    time.Now(),
// 		Ip:               lead.IP,
// 		Status:           LeadStatus.New,
// 		Stage:            LeadStage.S1,
// 		LeadStatusId:     1,
// 		QdeConsent:       "Y",
// 		CreatedOn:        time.Now(),
// 		TermAndCondition: lead.TermAndCondition,
// 		Coordinates:      lead.Coordinates,
// 		UtmSource:        lead.UtmSource,
// 		UtmCampaign:      lead.UtmCampaign,
// 		Promocode:        lead.Promocode,
// 	}
// }

// func (l *Lead) Insert(conn *db.Conn) error {
// 	conn.GetDB().Create(l)
// 	return nil
// }

// func GetLeadsByLeadId(conn *db.Conn, LeadId uint64) (*Lead, error) {
// 	var leads Lead
// 	err := conn.GetDB().Where("lead_active=1 AND lead_deleted=0 AND lead_id = ?", LeadId).Find(&leads).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &leads, nil
// }
