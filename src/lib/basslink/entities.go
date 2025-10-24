package basslink

import "encoding/json"

const (
	TableAdministrators           = "administrators"
	TableAdministratorCredentials = "administrator_credentials"
	TableAgents                   = "agents"
	TableAgentDocuments           = "agent_documents"
	TableAgentUsers               = "agent_users"
	TableAgentUserCredentials     = "agent_user_credentials"
	TableSenders                  = "senders"
	TableSenderDocuments          = "sender_documents"
	TableRecipients               = "recipients"
	TableRecipientDocuments       = "recipient_documents"
	TableRemittances              = "remittances"
	TableRemittanceAttachments    = "remittance_attachments"
	TableCurrencies               = "currencies"
	TableAppointments             = "appointments"
	TableRates                    = "rates"
	TableTemplates                = "templates"
	TableRemittancePayments       = "remittance_payments"

	AdministratorRoleRoot  = "root"
	AdministratorRoleSuper = "super"
	AdministratorRoleAdmin = "admin"

	AgentRoleOwner = "owner"
	AgentRoleAdmin = "admin"
	AgentRoleUser  = "user"
	AgentRoleGuest = "guest"

	RemittanceStatusSubmitted        = "submitted"         // submitted by user, waiting for approval
	RemittanceStatusWait             = "wait"              // approved by agent, waiting for payment
	RemittanceStatusPaymentConfirmed = "payment_confirmed" // payment confirmed by user, waiting for verification
	RemittanceStatusPaid             = "paid"              // payment approved by agent, waiting for transfer
	RemittanceStatusProcessed        = "processed"         // transfer being processed
	RemittanceStatusCompleted        = "completed"         // transfer completed
	RemittanceStatusCancelled        = "cancelled"         // not approved by agent, payment not received or expired
	RemittanceStatusFailed           = "refund"            // transfer failed

	PaymentStatusWait      = "wait"
	PaymentStatusCompleted = "completed"
	PaymentStatusFailed    = "failed"
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

type Sender struct {
	Id                string            `json:"id"`
	SenderType        string            `json:"sender_type"`
	Name              string            `json:"name"`
	Gender            string            `json:"gender"`
	Birthdate         string            `json:"birthdate"`
	Citizenship       string            `json:"citizenship"`
	IdentityType      string            `json:"identity_type"`
	IdentityNo        string            `json:"identity_no"`
	RegisteredCountry string            `json:"registered_country"`
	RegisteredRegion  string            `json:"registered_region"`
	RegisteredCity    string            `json:"registered_city"`
	RegisteredAddress string            `json:"registered_address"`
	RegisteredZipCode string            `json:"registered_zip_code"`
	Country           string            `json:"country"`
	Region            string            `json:"region"`
	City              string            `json:"city"`
	Address           string            `json:"address"`
	ZipCode           string            `json:"zip_code"`
	Contact           string            `json:"contact"`
	Occupation        string            `json:"occupation"`
	PepStatus         *string           `json:"pep_status,omitempty"`
	Notes             *string           `json:"notes,omitempty"`
	Created           int64             `json:"created"`
	CreatedBy         string            `json:"created_by"`
	Updated           *int64            `json:"updated,omitempty"`
	UpdatedBy         *string           `json:"updated_by,omitempty"`
	Documents         *[]SenderDocument `json:"documents,omitempty" gorm:"foreignKey:SenderId;reference:Id"`
}

func (t Sender) TableName() string { return TableSenders }

type SenderDocument struct {
	Id           string  `json:"id"`
	SenderId     string  `json:"sender_id"`
	DocumentType string  `json:"document_type"`
	DocumentData string  `json:"document_data"`
	Notes        *string `json:"notes,omitempty"`
	IsVerified   bool    `json:"is_verified"`
	Created      int64   `json:"created"`
	Updated      *int64  `json:"updated,omitempty"`
}

func (t SenderDocument) TableName() string { return TableSenderDocuments }

type Recipient struct {
	Id               string              `json:"id"`
	SenderId         string              `json:"sender_id"`
	RecipientType    string              `json:"recipient_type"`
	Relationship     string              `json:"relationship"`
	Name             string              `json:"name"`
	Country          string              `json:"country"`
	Region           string              `json:"region,"`
	City             string              `json:"city"`
	Address          string              `json:"address"`
	ZipCode          string              `json:"zip_code"`
	Contact          string              `json:"contact"`
	PepStatus        *string             `json:"pep_status,omitempty"`
	AccountType      string              `json:"account_type"`
	BankName         string              `json:"bank_name"`
	BankCode         string              `json:"bank_code"`
	BankAccountNo    string              `json:"bank_account_no"`
	BankAccountOwner string              `json:"bank_account_owner"`
	Notes            *string             `json:"notes,omitempty"`
	Created          int64               `json:"created"`
	Updated          *int64              `json:"updated,omitempty"`
	Documents        []RecipientDocument `json:"documents,omitempty" gorm:"foreignKey:RecipientId;reference:Id"`
}

func (t Recipient) TableName() string { return TableRecipients }

type RecipientDocument struct {
	Id           string  `json:"id"`
	RecipientId  string  `json:"recipient_id"`
	DocumentType string  `json:"document_type"`
	DocumentData string  `json:"document_data"`
	Notes        *string `json:"notes,omitempty"`
	IsVerified   bool    `json:"is_verified"`
	Created      int64   `json:"created"`
	Updated      *int64  `json:"updated,omitempty"`
}

func (t RecipientDocument) TableName() string { return TableRecipientDocuments }

type Remittance struct {
	Id                    string                  `json:"id"`
	AgentId               string                  `json:"agent_id"`
	SenderId              string                  `json:"sender_id"`
	FromCurrency          string                  `json:"from_currency"`
	FromAmount            float64                 `json:"from_amount"`
	FromSenderType        string                  `json:"from_sender_type"`
	FromName              string                  `json:"from_name"`
	FromGender            string                  `json:"from_gender"`
	FromBirthdate         string                  `json:"from_birthdate"`
	FromCitizenship       string                  `json:"from_citizenship"`
	FromIdentityType      string                  `json:"from_identity_type"`
	FromIdentityNo        string                  `json:"from_identity_no"`
	FromRegisteredCountry string                  `json:"from_registered_country"`
	FromRegisteredRegion  string                  `json:"from_registered_region"`
	FromRegisteredCity    string                  `json:"from_registered_city"`
	FromRegisteredAddress string                  `json:"from_registered_address"`
	FromRegisteredZipCode string                  `json:"from_registered_zip_code"`
	FromCountry           string                  `json:"from_country"`
	FromRegion            string                  `json:"from_region"`
	FromCity              string                  `json:"from_city"`
	FromAddress           string                  `json:"from_address"`
	FromZipCode           string                  `json:"from_zip_code"`
	FromContact           string                  `json:"from_contact"`
	FromOccupation        string                  `json:"from_occupation"`
	FromPepStatus         *string                 `json:"from_pep_status,omitempty"`
	FromNotes             *string                 `json:"from_notes"`
	RecipientId           string                  `json:"recipient_id"`
	ToCurrency            string                  `json:"to_currency"`
	ToAmount              float64                 `json:"to_amount"`
	ToRecipientType       string                  `json:"to_recipient_type"`
	ToRelationship        string                  `json:"to_relationship"`
	ToName                string                  `json:"to_name"`
	ToCountry             string                  `json:"to_country"`
	ToRegion              string                  `json:"to_region"`
	ToCity                string                  `json:"to_city"`
	ToAddress             string                  `json:"to_address"`
	ToZipCode             string                  `json:"to_zip_code"`
	ToContact             string                  `json:"to_contact"`
	ToPepStatus           *string                 `json:"to_pep_status,omitempty"`
	ToAccountType         string                  `json:"to_account_type,omitempty"`
	ToBankName            string                  `json:"to_bank_name"`
	ToBankCode            string                  `json:"to_bank_code"`
	ToBankAccountNo       string                  `json:"to_bank_account_no"`
	ToBankAccountOwner    string                  `json:"to_bank_account_owner"`
	ToNotes               *string                 `json:"to_notes,omitempty"`
	RateCurrency          string                  `json:"rate_currency"`
	Rate                  float64                 `json:"rate"`
	FeeCurrency           string                  `json:"fee_currency"`
	FeeAmountPercent      float64                 `json:"fee_amount_percent"`
	FeeAmountFixed        float64                 `json:"fee_amount_fixed"`
	FeeTotal              float64                 `json:"fee_total"`
	PaymentMethod         string                  `json:"payment_method"`
	TransferType          string                  `json:"transfer_type"`
	TransferRef           *string                 `json:"transfer_ref,omitempty"`
	FundSource            *string                 `json:"fund_source,omitempty"`
	Purpose               *string                 `json:"purpose,omitempty"`
	Notes                 *string                 `json:"notes,omitempty"`
	Status                string                  `json:"status"`
	IsSettled             bool                    `json:"is_settled"`
	CreatedBy             *string                 `json:"created_by"`
	ApprovedBy            *string                 `json:"approved_by"`
	ReleasedBy            *string                 `json:"released_by"`
	ApprovedAt            *int64                  `json:"approved_at,omitempty"`
	ReleasedAt            *int64                  `json:"released_at,omitempty"`
	Created               int64                   `json:"created"`
	Updated               *int64                  `json:"updated,omitempty"`
	SourceCurrency        *Currency               `json:"source_currency,omitempty" gorm:"foreignKey:FromCurrency;reference:Id"`
	TargetCurrency        *Currency               `json:"target_currency,omitempty" gorm:"foreignKey:ToCurrency;reference:Id"`
	Attachments           *[]RemittanceAttachment `json:"attachments" gorm:"foreignKey:RemittanceId;reference:Id"`
}

func (t Remittance) TableName() string { return TableRemittances }

type RemittanceAttachment struct {
	Id           string `json:"id"`
	RemittanceId string `json:"remittance_id"`
	Attachment   string `json:"attachment"`
	SubmitBy     string `json:"submit_by"`
	SubmitTime   int64  `json:"submit_time"`
}

func (t RemittanceAttachment) TableName() string { return TableRemittanceAttachments }

type Currency struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	AllowIn  bool   `json:"allow_in"`
	AllowOut bool   `json:"allow_out"`
	Decimal  int8   `json:"decimal"`
}

func (t Currency) TableName() string { return TableCurrencies }

type Appointment struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Company     *string `json:"company,omitempty"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	ServiceType string  `json:"service_type"`
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	Notes       *string `json:"notes,omitempty"`
	Status      string  `json:"status"`
	Created     int64   `json:"created"`
	Updated     *int64  `json:"updated,omitempty"`
}

func (t Appointment) TableName() string { return TableAppointments }

type Rate struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Rate         float64 `json:"rate"`
	Source       *string `json:"source,omitempty"`
	Updated      int64   `json:"updated"`
}

func (t Rate) TableName() string { return TableRates }

type Template struct {
	Id           string `json:"id"`
	TemplateType string `json:"template_type"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	Data         string `json:"data"`
	Created      int64  `json:"created"`
	Updated      *int64 `json:"updated"`
}

func (t Template) TableName() string { return TableTemplates }

type RemittancePayment struct {
	Id                  string      `json:"id"`
	Currency            string      `json:"currency"`
	Amount              float64     `json:"amount"`
	PaymentMethod       string      `json:"payment_method"`
	PaymentData         string      `json:"payment_data"`
	PaymentConfirmTime  *string     `json:"payment_confirm_time"`
	PaymentConfirmProof *string     `json:"payment_confirm_proof"`
	PaymentReference    *string     `json:"payment_reference"`
	Status              string      `json:"status"`
	Created             int64       `json:"created"`
	Updated             *int64      `json:"updated"`
	Remittance          *Remittance `json:"remittance,omitempty" gorm:"foreignKey:Id;reference:Id"`
}

func (t RemittancePayment) TableName() string { return TableRemittancePayments }

type RateInfo struct {
	FromCurrency string      `json:"from_currency"`
	ToCurrency   string      `json:"to_currency"`
	FromAmount   json.Number `json:"from_amount"`
	ToAmount     json.Number `json:"to_amount"`
	Rate         json.Number `json:"rate"`
	FeePercent   json.Number `json:"fee_percent"`
	FeeFixed     json.Number `json:"fee_fixed"`
	TotalFee     json.Number `json:"total_fee"`
}

type BankInfo struct {
	BankName     string `json:"bank_name"`
	BankCode     string `json:"bank_code"`
	SwiftCode    string `json:"swift_code"`
	AccountNo    string `json:"account_no"`
	AccountOwner string `json:"account_owner"`
	Currency     string `json:"currency"`
}
