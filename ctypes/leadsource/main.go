package leadsource

import "salarykart/fmvalidator"

type LeadSource string

const (
	WebsiteFD       LeadSource = "Website FD"
	Appsalarykart    LeadSource = "Appsalarykart"
	Import          LeadSource = "Import"
	C4C             LeadSource = "C4C"
	APPFDIOS        LeadSource = "APPFDIOS"
	REFCASE         LeadSource = "REFCASE"
	Chatbot         LeadSource = "Chatbot"
	WebsiteFM       LeadSource = "Website FM"
	salarykartIn     LeadSource = "salarykart.IN"
	AppIOS          LeadSource = "AppIOS"
	WebsiteFundsMam LeadSource = "Website FundsMam"
)

func Validator() fmvalidator.Validator {
	var leadSources = []LeadSource{
		WebsiteFD,
		Appsalarykart,
		Import,
		C4C,
		APPFDIOS,
		REFCASE,
		Chatbot,
		WebsiteFM,
		salarykartIn,
		AppIOS,
		WebsiteFundsMam,
	}

	var s []string
	for _, v := range leadSources {
		s = append(s, string(v))
	}
	return fmvalidator.NewValidator("lead_source", s)
}
