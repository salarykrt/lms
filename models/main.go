package models

import (
	"salarykart/db"
	masters "salarykart/models/masters"
)

func MigrateAll(c *db.Conn) error {
	err := c.GetDB().AutoMigrate(
		&Lead{},
		&LeadFollowup{},
		&LeadsOtpTran{},
		&CustomerDetail{},
		&CustomerBanking{},
		&CifCustomer{},
		&CifCustomer{},
		&CreditAnalysisMemo{},
		&Collection{},
		&Document{},
		&CompanyHoliday{},
		&CompanyLogin{},
		&CustomerEmployment{},
		&CustomerReferences{},

		// Masters
		&masters.MasterRelationType{},
		&masters.MasterFoir{},
		&masters.MasterDocument{},
		&masters.MasterBankType{},
		&masters.MasterBankAccountStatus{},
		&masters.MasterCity{},
		&masters.MasterCompanyAccounts{},
		&masters.MasterCompanyAccounts{},
	)
	return err
}
