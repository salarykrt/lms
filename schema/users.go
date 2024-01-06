package schema

import "time"

type Users struct {
	UserID                        uint64    `json:"user_id" binding:"omitempty"`
	UserName                      string    `json:"user_name" binding:"omitempty"`
	RoleID                        uint      `json:"role_id" binding:"omitempty"`
	Labels                        string    `json:"labels" binding:"omitempty"`
	Role                          string    `json:"role" binding:"omitempty"`
	Name                          string    `json:"name" binding:"omitempty"`
	Email                         string    `json:"email" binding:"omitempty"`
	Mobile                        uint64    `json:"mobile" binding:"omitempty"`
	Password                      string    `json:"password" binding:"omitempty"`
	Dob                           string    `json:"dob" binding:"omitempty"`
	MaritalStatus                 string    `json:"marital_status" binding:"omitempty"`
	FatherName                    string    `json:"father_name" binding:"omitempty"`
	Branch                        string    `json:"branch" binding:"omitempty"`
	Center                        string    `json:"center" binding:"omitempty"`
	Status                        uint8     `json:"status" binding:"omitempty"`
	Otp                           int       `json:"otp" binding:"omitempty"`
	LastActivity                  string    `json:"last_activity" binding:"omitempty"`
	IsActive                      int       `json:"is_Active" binding:"omitempty"`
	IP                            string    `json:"ip" binding:"omitempty"`
	CreatedBy                     uint      `json:"created_by" binding:"omitempty"`
	CreatedOn                     time.Time `json:"created_on" binding:"omitempty"`
	UserScmID                     uint      `json:"user_scm_id" binding:"omitempty"`
	UserActive                    uint      `json:"user_active" binding:"omitempty"`
	UserDeleted                   uint      `json:"user_deleted" binding:"omitempty"`
	UpdatedBy                     uint      `json:"updated_by" binding:"omitempty"`
	UpdatedOn                     time.Time `json:"updated_on" binding:"omitempty"`
	UserLastLoginDatetime         time.Time `json:"user_last_login_datetime" binding:"omitempty"`
	UserDialerID                  string    `json:"user_dialer_id" binding:"omitempty"`
	QcAgentCode                   uint64    `json:"qc_agent_code" binding:"omitempty"`
	DiallerCampaignID             uint      `json:"dialler_campaign_id" binding:"omitempty"`
	UserStatusID                  uint      `json:"user_status_id" binding:"omitempty"`
	UserIsLoanwalle               uint      `json:"user_is_loanwalle" binding:"omitempty"`
	UserLastLoginIP               string    `json:"user_last_login_ip" binding:"omitempty"`
	ValidIPRequired               uint      `json:"valid_ip_required" binding:"omitempty"`
	UserTotalLoginCount           uint      `json:"user_total_login_count" binding:"omitempty"`
	UserLoginsFailedCount         int       `json:"user_logins_failed_count" binding:"omitempty"`
	UserLastPasswordResetDatetime time.Time `json:"user_last_password_reset_datetime" binding:"omitempty"`
}
