package schema

import (
	"salarykart/ctypes/leadsource"

	"cloud.google.com/go/civil"
)

type Lead struct {
	LoanAmount       int                   `json:"loan_amount" binding:"required,gt=0"`
	Name             string                `json:"name" binding:"required"`
	Gender           string                `json:"gender" binding:"required"`
	DateOfBirth      civil.Date            `json:"date_of_birth" binding:"required"`
	PanCard          string                `json:"pan_card" binding:"required" min_length:"10" max_length:"10"`
	Mobile           string                `json:"mobile" binding:"required,numeric,min=10,max=10"`
	AlternateMobile  string                `json:"alternate_mobile" binding:"omitempty,numeric,min=10,max=10"`
	Obligations      int                   `json:"obligations" binding:"omitempty,numeric"`
	EmailPersonal    string                `json:"email_personal" binding:"omitempty,email"`
	EmailOffice      string                `json:"email_office" binding:"omitempty,email"`
	Source           leadsource.LeadSource `json:"source" binding:"required,lead_source"`
	Coordinates      string                `json:"coordinates" binding:"omitempty"`
	IP               string                `json:"ip" binding:"omitempty"`
	StateId          uint32                `json:"state_id" binding:"required"`
	CityId           uint64                `json:"city_id" binding:"required"`
	Pincode          string                `json:"pincode" binding:"required,numeric,min=6,max=6"`
	SourceId         uint32                `json:"source_id" binding:"omitempty"`
	UserType         string                `json:"user_type" binding:"omitempty"`
	TermAndCondition string                `json:"term_and_condition" binding:"omitempty"`
	UtmSource        string                `json:"utm_source" binding:"omitempty"`
	UtmCampaign      string                `json:"utm_campaign" binding:"omitempty"`
	Promocode        string                `json:"promocode" binding:"omitempty"`
}
