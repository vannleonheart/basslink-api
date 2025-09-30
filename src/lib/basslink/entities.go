package basslink

const (
	TableAdministrators           = "administrators"
	TableAdministratorCredentials = "administrator_credentials"
	TableAgents                   = "agents"
	TableAgentDocuments           = "agent_documents"
	TableAgentUsers               = "agent_users"
	TableAgentUserCredentials     = "agent_user_credentials"
	TableUsers                    = "users"
	TableUserCredentials          = "user_credentials"
	TableUserDocuments            = "user_documents"
	TableContacts                 = "contacts"
	TableContactDocuments         = "contact_documents"
	TableContactAccounts          = "contact_accounts"
	TableDeposits                 = "deposits"
	TableDisbursements            = "disbursements"
	TableCurrencies               = "currencies"

	AdministratorRoleRoot  = "root"
	AdministratorRoleSuper = "super"
	AdministratorRoleAdmin = "admin"

	AgentRoleOwner = "owner"

	DisbursementStatusDraft     = "draft"
	DisbursementStatusNew       = "new"
	DisbursementStatusApproved  = "approved"
	DisbursementStatusRejected  = "rejected"
	DisbursementStatusPending   = "pending"
	DisbursementStatusProcessed = "processed"
	DisbursementStatusCompleted = "paid"
	DisbursementStatusCancelled = "cancelled"
	DisbursementStatusFailed    = "failed"
)

type Administrator struct {
	Id        string  `json:"id"`
	Role      string  `json:"role"`
	Username  string  `json:"username"`
	Name      string  `json:"name"`
	Email     *string `json:"email,omitempty"`
	PhoneCode *string `json:"phone_code,omitempty"`
	PhoneNo   *string `json:"phone_no,omitempty"`
	IsEnable  bool    `json:"is_enable"`
	Created   int64   `json:"created"`
	Updated   *int64  `json:"updated"`
}

func (t Administrator) TableName() string { return TableAdministrators }

type AdministratorCredential struct {
	AdministratorId string `json:"administrator_id"`
	CredentialType  string `json:"credential_type"`
	CredentialData  string `json:"credential_data"`
	Updated         int64  `json:"updated,omitempty"`
}

func (t AdministratorCredential) TableName() string { return TableAdministratorCredentials }

type Agent struct {
	Id         string  `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Country    *string `json:"country,omitempty"`
	Region     *string `json:"region,omitempty"`
	City       *string `json:"city,omitempty"`
	Address    *string `json:"address,omitempty"`
	PhoneCode  *string `json:"phone_code,omitempty"`
	PhoneNo    *string `json:"phone_no,omitempty"`
	Email      *string `json:"email,omitempty"`
	Website    *string `json:"website,omitempty"`
	Timezone   *int    `json:"timezone,omitempty"`
	IsVerified bool    `json:"is_verified"`
	IsEnable   bool    `json:"is_enabled"`
	Created    int64   `json:"created"`
	Updated    *int64  `json:"updated,omitempty"`
}

func (t Agent) TableName() string { return TableAgents }

type AgentDocument struct {
	Id           string  `json:"id"`
	AgentId      string  `json:"agent_id"`
	DocumentType string  `json:"document_type"`
	DocumentData string  `json:"document_data"`
	Notes        *string `json:"notes,omitempty"`
	IsVerified   bool    `json:"is_verified"`
	Created      int64   `json:"created"`
	Updated      *int64  `json:"updated,omitempty"`
}

func (t AgentDocument) TableName() string { return TableAgentDocuments }

type AgentUser struct {
	Id        string  `json:"id"`
	AgentId   string  `json:"agent_id"`
	Role      string  `json:"role"`
	Username  string  `json:"username"`
	Name      string  `json:"name"`
	Email     *string `json:"email,omitempty"`
	PhoneCode *string `json:"phone_code,omitempty"`
	PhoneNo   *string `json:"phone_no,omitempty"`
	IsEnable  bool    `json:"is_enable"`
	Created   int64   `json:"created"`
	Updated   *int64  `json:"updated,omitempty"`
	Agent     *Agent  `json:"agent,omitempty" gorm:"foreignKey:AgentId;reference:Id"`
}

func (t AgentUser) TableName() string { return TableAgentUsers }

type AgentUserCredential struct {
	AgentUserId    string `json:"agent_user_id"`
	CredentialType string `json:"credential_type"`
	CredentialData string `json:"credential_data"`
	Updated        int64  `json:"updated,omitempty"`
}

func (t AgentUserCredential) TableName() string { return TableAgentUserCredentials }

type User struct {
	Id            string  `json:"id"`
	AgentId       string  `json:"agent_id"`
	Username      *string `json:"username"`
	UserType      string  `json:"user_type"`
	Name          string  `json:"name"`
	Gender        *string `json:"gender,omitempty"`
	Birthdate     *string `json:"birthdate,omitempty"`
	Citizenship   string  `json:"citizenship"`
	IdentityType  string  `json:"identity_type"`
	IdentityNo    string  `json:"identity_no"`
	Country       *string `json:"country,omitempty"`
	Region        *string `json:"region,omitempty"`
	City          *string `json:"city,omitempty"`
	Address       *string `json:"address,omitempty"`
	Email         *string `json:"email,omitempty"`
	PhoneCode     *string `json:"phone_code,omitempty"`
	PhoneNo       *string `json:"phone_no,omitempty"`
	Occupation    *string `json:"occupation,omitempty"`
	Notes         *string `json:"notes,omitempty"`
	IsVerified    bool    `json:"is_verified"`
	EmailVerified bool    `json:"email_verified"`
	PhoneVerified bool    `json:"phone_verified"`
	IsEnable      bool    `json:"is_enabled"`
	Created       int64   `json:"created,omitempty"`
	Updated       *int64  `json:"updated,omitempty"`
}

func (t User) TableName() string { return TableUsers }

type UserCredential struct {
	UserId         string `json:"user_id"`
	CredentialType string `json:"credential_type"`
	CredentialData string `json:"credential_data"`
	Updated        int64  `json:"updated"`
}

func (t UserCredential) TableName() string { return TableUserCredentials }

type UserDocument struct {
	Id           string  `json:"id"`
	UserId       string  `json:"user_id"`
	DocumentType string  `json:"document_type"`
	DocumentData string  `json:"document_data"`
	Notes        *string `json:"notes,omitempty"`
	IsVerified   bool    `json:"is_verified"`
	Created      int64   `json:"created"`
	Updated      *int64  `json:"updated,omitempty"`
}

func (t UserDocument) TableName() string { return TableUserDocuments }

type Contact struct {
	Id           string            `json:"id"`
	AgentId      string            `json:"agent_id"`
	ContactType  string            `json:"contact_type"`
	Name         string            `json:"name"`
	Gender       *string           `json:"gender,omitempty"`
	Birthdate    *string           `json:"birthdate,omitempty"`
	Citizenship  string            `json:"citizenship"`
	IdentityType string            `json:"identity_type"`
	IdentityNo   string            `json:"identity_no"`
	Country      *string           `json:"country,omitempty"`
	Region       *string           `json:"region,omitempty"`
	City         *string           `json:"city,omitempty"`
	Address      *string           `json:"address,omitempty"`
	Email        *string           `json:"email,omitempty"`
	PhoneCode    *string           `json:"phone_code,omitempty"`
	PhoneNo      *string           `json:"phone_no,omitempty"`
	Occupation   *string           `json:"occupation,omitempty"`
	Notes        *string           `json:"notes,omitempty"`
	IsVerified   bool              `json:"is_verified"`
	Created      int64             `json:"created"`
	Updated      *int64            `json:"updated,omitempty"`
	Documents    []ContactDocument `json:"documents,omitempty" gorm:"foreignKey:ContactId;reference:Id"`
	Accounts     []ContactAccount  `json:"accounts,omitempty" gorm:"foreignKey:ContactId;reference:Id"`
}

func (t Contact) TableName() string { return TableContacts }

type ContactDocument struct {
	Id           string  `json:"id"`
	ContactId    string  `json:"contact_id"`
	DocumentType string  `json:"document_type"`
	DocumentData string  `json:"document_data"`
	Notes        *string `json:"notes,omitempty"`
	IsVerified   bool    `json:"is_verified"`
	Created      int64   `json:"created"`
	Updated      *int64  `json:"updated,omitempty"`
}

func (t ContactDocument) TableName() string { return TableContactDocuments }

type ContactAccount struct {
	Id          string  `json:"id"`
	ContactId   string  `json:"contact_id"`
	AccountType string  `json:"account_type"`
	BankId      string  `json:"bank_id"`
	BankName    string  `json:"bank_name"`
	BankCode    *string `json:"bank_code,omitempty"`
	BankSwift   *string `json:"bank_swift,omitempty"`
	OwnerName   string  `json:"owner_name"`
	No          string  `json:"no"`
	Country     *string `json:"country,omitempty"`
	Address     *string `json:"address,omitempty"`
	Email       *string `json:"email,omitempty"`
	Website     *string `json:"website,omitempty"`
	PhoneCode   *string `json:"phone_code,omitempty"`
	PhoneNo     *string `json:"phone_no,omitempty"`
	Notes       *string `json:"notes,omitempty"`
	Created     int64   `json:"created"`
	Updated     *int64  `json:"updated,omitempty"`
}

func (t ContactAccount) TableName() string { return TableContactAccounts }

type Disbursement struct {
	Id             string          `json:"id"`
	AgentId        string          `json:"agent_id"`
	UserId         *string         `json:"user_id,omitempty"`
	FromCurrency   string          `json:"from_currency"`
	FromAmount     float64         `json:"from_amount"`
	ToContact      string          `json:"to_contact"`
	ToCurrency     string          `json:"to_currency"`
	ToAmount       float64         `json:"to_amount"`
	ToAccount      string          `json:"to_account"`
	RateCurrency   string          `json:"rate_currency"`
	Rate           float64         `json:"rate"`
	FeeCurrency    string          `json:"fee_currency"`
	FeeAmount      float64         `json:"fee_amount"`
	TransferType   string          `json:"transfer_type"`
	TransferRef    *string         `json:"transfer_ref,omitempty"`
	TransferDate   *string         `json:"transfer_date,omitempty"`
	FundSource     *string         `json:"fund_source,omitempty"`
	Purpose        *string         `json:"purpose,omitempty"`
	Notes          *string         `json:"notes,omitempty"`
	Status         string          `json:"status"`
	IsSettled      bool            `json:"is_settled"`
	Created        int64           `json:"created"`
	Updated        *int64          `json:"updated,omitempty"`
	TargetAccount  *ContactAccount `json:"target_account,omitempty" gorm:"foreignKey:ToAccount;reference:Id"`
	TargetCurrency *Currency       `json:"target_currency,omitempty" gorm:"foreignKey:ToCurrency;reference:Id"`
	User           *User           `json:"user,omitempty" gorm:"foreignKey:UserId;reference:Id"`
	Contact        *Contact        `json:"contact,omitempty" gorm:"foreignKey:ToContact;reference:Id"`
}

func (t Disbursement) TableName() string { return TableDisbursements }

type Currency struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	IsActive bool   `json:"is_active"`
}

func (t Currency) TableName() string { return TableCurrencies }
