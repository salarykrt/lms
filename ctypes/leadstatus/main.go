package leadstatus

type LeadStatus string

const (
	New                    LeadStatus = "LEAD-NEW"
	LeadInProcess          LeadStatus = "LEAD-INPROCESS"
	LeadHold               LeadStatus = "LEAD-HOLD"
	ApplicationNew         LeadStatus = "APPLICATION-NEW"
	ApplicationInProcess   LeadStatus = "APPLICATION-INPROCESS"
	ApplicationHold        LeadStatus = "APPLICATION-HOLD"
	Duplicate              LeadStatus = "DUPLICATE"
	Reject                 LeadStatus = "REJECT"
	ApplicationRecommended LeadStatus = "APPLICATION-RECOMMENDED"
	ApplicationSendBack    LeadStatus = "APPLICATION-SEND-BACK"
	Sanction               LeadStatus = "SANCTION"
	DisbursePending        LeadStatus = "DISBURSE-PENDING"
	Disbursed              LeadStatus = "DISBURSED"
	Cancel                 LeadStatus = "CANCEL"
	PartPayment            LeadStatus = "PART-PAYMENT"
	Closed                 LeadStatus = "CLOSED"
	Settled                LeadStatus = "SETTLED"
	WriteOff               LeadStatus = "WRITEOFF"
	DisbursalNew           LeadStatus = "DISBURSAL-NEW"
	DisbursalInProcess     LeadStatus = "DISBURSAL-INPROCESS"
	DisbursalHold          LeadStatus = "DISBURSAL-HOLD"
	DisbursedWaived        LeadStatus = "DISBURSED-WAIVED"
	DisbursalSendBack      LeadStatus = "DISBURSAL-SEND-BACK"
)
