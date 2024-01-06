package leadstage

import "salarykart/fmvalidator"

type LeadStage string

const (
	S1  LeadStage = "S1"
	S2  LeadStage = "S2"
	S3  LeadStage = "S3"
	S4  LeadStage = "S4"
	S5  LeadStage = "S5"
	S6  LeadStage = "S6"
	S7  LeadStage = "S7"
	S8  LeadStage = "S8"
	S9  LeadStage = "S9"
	S10 LeadStage = "S10"
	S11 LeadStage = "S11"
	S12 LeadStage = "S12"
	S13 LeadStage = "S13"
	S14 LeadStage = "S14"
	S15 LeadStage = "S15"
	S16 LeadStage = "S16"
	S17 LeadStage = "S17"
	S18 LeadStage = "S18"
	S19 LeadStage = "S19"
	S20 LeadStage = "S20"
	S21 LeadStage = "S21"
	S22 LeadStage = "S22"
	S25 LeadStage = "S25"
	S30 LeadStage = "S30"
)

func Validator() fmvalidator.Validator {
	var LeadStages = []LeadStage{
		S1,
		S2,
		S3,
		S4,
		S5,
		S6,
		S7,
		S8,
		S9,
		S10,
		S11,
		S12,
		S13,
		S14,
		S15,
		S16,
		S17,
		S18,
		S19,
		S20,
		S21,
		S22,
		S25,
		S30,
	}

	var s []string
	for _, v := range LeadStages {
		s = append(s, string(v))
	}
	return fmvalidator.NewValidator("lead_stage", s)
}
