package agent

import "encoding/json"

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type UpdatePasswordRequest struct {
	Password                string `json:"password" validate:"required"`
	NewPassword             string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,eqfield=NewPassword"`
}

type CreateAgentUserRequest struct {
	Role                 string  `json:"role" validate:"required"`
	Username             string  `json:"username" validate:"required,min=5,max=50"`
	Name                 string  `json:"name" validate:"required"`
	Email                *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode            *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo              *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Password             string  `json:"password" validate:"required,min=8"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type UpdateAgentUserRequest struct {
	Role                 string  `json:"role" validate:"required"`
	Username             string  `json:"username" validate:"required,min=5,max=50"`
	Name                 string  `json:"name" validate:"required"`
	Email                *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode            *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo              *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Password             *string `json:"password,omitempty" validate:"omitempty,min=8"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty" validate:"omitempty,eqfield=Password"`
}

type CreateSenderRequest struct {
	SenderType              string  `json:"sender_type" validate:"required,oneof=individual institution"`
	SenderName              string  `json:"sender_name" validate:"required,min=5,max=100"`
	SenderGender            *string `json:"sender_gender,omitempty" validate:"omitempty,oneof=male female"`
	SenderBirthdate         *string `json:"sender_birthdate,omitempty" validate:"omitempty"`
	SenderCitizenship       string  `json:"sender_citizenship" validate:"required"`
	SenderRegisteredCountry string  `json:"sender_registered_country" validate:"required"`
	SenderRegisteredRegion  string  `json:"sender_registered_region" validate:"required"`
	SenderRegisteredCity    string  `json:"sender_registered_city" validate:"required"`
	SenderRegisteredAddress string  `json:"sender_registered_address" validate:"required"`
	SenderRegisteredZipCode string  `json:"sender_registered_zip_code" validate:"required"`
	SenderCountry           string  `json:"sender_country" validate:"required"`
	SenderRegion            string  `json:"sender_region" validate:"required"`
	SenderCity              string  `json:"sender_city" validate:"required"`
	SenderAddress           string  `json:"sender_address" validate:"required"`
	SenderZipCode           string  `json:"sender_zip_code" validate:"required"`
	SenderContact           string  `json:"sender_contact" validate:"required"`
	SenderIdentityType      string  `json:"sender_identity_type,omitempty" validate:"required"`
	SenderIdentityNo        string  `json:"sender_identity_no,omitempty" validate:"required,min=3,max=100"`
	SenderOccupation        string  `json:"sender_occupation,omitempty" validate:"omitempty"`
	SenderPepStatus         *string `json:"sender_pep_status,omitempty" validate:"omitempty"`
	SenderNotes             *string `json:"sender_notes,omitempty" validate:"omitempty"`
	SenderDocuments         *[]struct {
		DocumentData *string `json:"document_data" validate:"omitempty"`
		DocumentType *string `json:"document_type" validate:"omitempty"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"sender_documents" validate:"omitempty,dive"`
}

type UpdateSenderRequest struct {
	SenderType              string  `json:"sender_type" validate:"required,oneof=individual institution"`
	SenderName              string  `json:"sender_name" validate:"required,min=5,max=100"`
	SenderGender            *string `json:"sender_gender,omitempty" validate:"omitempty,oneof=male female"`
	SenderBirthdate         *string `json:"sender_birthdate,omitempty" validate:"omitempty"`
	SenderCitizenship       string  `json:"sender_citizenship" validate:"required"`
	SenderRegisteredCountry string  `json:"sender_registered_country" validate:"required"`
	SenderRegisteredRegion  string  `json:"sender_registered_region" validate:"required"`
	SenderRegisteredCity    string  `json:"sender_registered_city" validate:"required"`
	SenderRegisteredAddress string  `json:"sender_registered_address" validate:"required"`
	SenderRegisteredZipCode string  `json:"sender_registered_zip_code" validate:"required"`
	SenderCountry           string  `json:"sender_country" validate:"required"`
	SenderRegion            string  `json:"sender_region" validate:"required"`
	SenderCity              string  `json:"sender_city" validate:"required"`
	SenderAddress           string  `json:"sender_address" validate:"required"`
	SenderZipCode           string  `json:"sender_zip_code" validate:"required"`
	SenderContact           string  `json:"sender_contact" validate:"required"`
	SenderIdentityType      string  `json:"sender_identity_type,omitempty" validate:"required"`
	SenderIdentityNo        string  `json:"sender_identity_no,omitempty" validate:"required,min=3,max=100"`
	SenderOccupation        string  `json:"sender_occupation,omitempty" validate:"omitempty"`
	SenderPepStatus         *string `json:"sender_pep_status,omitempty" validate:"omitempty"`
	SenderNotes             *string `json:"sender_notes,omitempty" validate:"omitempty"`
	SenderDocuments         *[]struct {
		Id           *string `json:"id,omitempty" validate:"omitempty"`
		DocumentData *string `json:"document_data" validate:"omitempty"`
		DocumentType *string `json:"document_type" validate:"omitempty"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"sender_documents" validate:"omitempty,dive"`
}

type CreateRecipientRequest struct {
	RecipientSenderId         string  `json:"recipient_sender_id" validate:"required"`
	RecipientType             string  `json:"recipient_type" validate:"required,oneof=individual institution"`
	RecipientRelationship     string  `json:"recipient_relationship" validate:"required"`
	RecipientName             string  `json:"recipient_name" validate:"required"`
	RecipientCountry          string  `json:"recipient_country" validate:"required"`
	RecipientRegion           string  `json:"recipient_region" validate:"required"`
	RecipientCity             string  `json:"recipient_city" validate:"required"`
	RecipientAddress          string  `json:"recipient_address" validate:"required"`
	RecipientZipCode          string  `json:"recipient_zip_code" validate:"required"`
	RecipientContact          string  `json:"recipient_contact" validate:"required"`
	RecipientPepStatus        *string `json:"recipient_pep_status,omitempty" validate:"omitempty"`
	RecipientAccountType      string  `json:"recipient_account_type" validate:"required"`
	RecipientBankName         string  `json:"recipient_bank_name" validate:"required"`
	RecipientBankCode         string  `json:"recipient_bank_code" validate:"required"`
	RecipientBankAccountNo    string  `json:"recipient_bank_account_no" validate:"required"`
	RecipientBankAccountOwner string  `json:"recipient_bank_account_owner" validate:"required"`
	RecipientNotes            *string `json:"recipient_notes,omitempty" validate:"omitempty"`
	RecipientDocuments        *[]struct {
		DocumentData *string `json:"document_data" validate:"omitempty"`
		DocumentType *string `json:"document_type" validate:"omitempty"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"recipient_documents" validate:"omitempty,dive"`
}

type UpdateRecipientRequest struct {
	RecipientType             string  `json:"recipient_type" validate:"required,oneof=individual institution"`
	RecipientRelationship     string  `json:"recipient_relationship" validate:"required"`
	RecipientName             string  `json:"recipient_name" validate:"required"`
	RecipientCountry          string  `json:"recipient_country" validate:"required"`
	RecipientRegion           string  `json:"recipient_region" validate:"required"`
	RecipientCity             string  `json:"recipient_city" validate:"required"`
	RecipientAddress          string  `json:"recipient_address" validate:"required"`
	RecipientZipCode          string  `json:"recipient_zip_code" validate:"required"`
	RecipientContact          string  `json:"recipient_contact" validate:"required"`
	RecipientPepStatus        *string `json:"recipient_pep_status,omitempty" validate:"omitempty"`
	RecipientAccountType      string  `json:"recipient_account_type" validate:"required"`
	RecipientBankName         string  `json:"recipient_bank_name" validate:"required"`
	RecipientBankCode         string  `json:"recipient_bank_code" validate:"required"`
	RecipientBankAccountNo    string  `json:"recipient_bank_account_no" validate:"required"`
	RecipientBankAccountOwner string  `json:"recipient_bank_account_owner" validate:"required"`
	RecipientNotes            *string `json:"recipient_notes,omitempty" validate:"omitempty"`
	RecipientDocuments        *[]struct {
		Id           *string `json:"id,omitempty" validate:"omitempty"`
		DocumentData *string `json:"document_data" validate:"omitempty"`
		DocumentType *string `json:"document_type" validate:"omitempty"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"recipient_documents" validate:"omitempty,dive"`
}

type CreateRecipientDocumentRequest struct {
	DocumentType string  `json:"document_type" validate:"required"`
	DocumentData string  `json:"document_data" validate:"required"`
	Notes        *string `json:"notes,omitempty" validate:"omitempty"`
}

type UpdateRecipientDocumentRequest struct {
	DocumentType string  `json:"document_type" validate:"required"`
	DocumentData string  `json:"document_data" validate:"required"`
	Notes        *string `json:"notes,omitempty" validate:"omitempty"`
}

type CreateRemittanceRequest struct {
	FromCurrency              string      `json:"from_currency" validate:"required"`
	FromAmount                json.Number `json:"from_amount" validate:"required"`
	SenderId                  *string     `json:"sender_id,omitempty" validate:"omitempty"`
	SenderType                string      `json:"sender_type" validate:"required,oneof=individual institution"`
	SenderName                string      `json:"sender_name" validate:"required,min=5,max=100"`
	SenderGender              *string     `json:"sender_gender,omitempty" validate:"omitempty,oneof=male female"`
	SenderBirthdate           *string     `json:"sender_birthdate,omitempty" validate:"omitempty"`
	SenderCitizenship         string      `json:"sender_citizenship" validate:"required"`
	SenderRegisteredCountry   string      `json:"sender_registered_country" validate:"required"`
	SenderRegisteredRegion    string      `json:"sender_registered_region" validate:"required"`
	SenderRegisteredCity      string      `json:"sender_registered_city" validate:"required"`
	SenderRegisteredAddress   string      `json:"sender_registered_address" validate:"required"`
	SenderRegisteredZipCode   string      `json:"sender_registered_zip_code" validate:"required"`
	SenderCountry             string      `json:"sender_country" validate:"required"`
	SenderRegion              string      `json:"sender_region" validate:"required"`
	SenderCity                string      `json:"sender_city" validate:"required"`
	SenderAddress             string      `json:"sender_address" validate:"required"`
	SenderZipCode             string      `json:"sender_zip_code" validate:"required"`
	SenderContact             string      `json:"sender_contact" validate:"required"`
	SenderIdentityType        string      `json:"sender_identity_type" validate:"required"`
	SenderIdentityNo          string      `json:"sender_identity_no" validate:"required,min=3,max=100"`
	SenderOccupation          string      `json:"sender_occupation" validate:"omitempty"`
	SenderPepStatus           *string     `json:"sender_pep_status,omitempty" validate:"omitempty"`
	SenderNotes               *string     `json:"sender_notes,omitempty" validate:"omitempty"`
	SenderUpdate              *bool       `json:"sender_update,omitempty" validate:"omitempty"`
	ToCurrency                string      `json:"to_currency,omitempty" validate:"required"`
	ToAmount                  json.Number `json:"to_amount" validate:"required"`
	RecipientId               *string     `json:"recipient_id,omitempty" validate:"omitempty"`
	RecipientType             string      `json:"recipient_type" validate:"required,oneof=individual institution"`
	RecipientRelationship     string      `json:"recipient_relationship" validate:"required"`
	RecipientName             string      `json:"recipient_name" validate:"required"`
	RecipientCountry          string      `json:"recipient_country" validate:"required"`
	RecipientRegion           string      `json:"recipient_region" validate:"required"`
	RecipientCity             string      `json:"recipient_city" validate:"required"`
	RecipientAddress          string      `json:"recipient_address" validate:"required"`
	RecipientZipCode          string      `json:"recipient_zip_code" validate:"required"`
	RecipientContact          string      `json:"recipient_contact" validate:"required"`
	RecipientPepStatus        *string     `json:"recipient_pep_status,omitempty" validate:"omitempty"`
	RecipientAccountType      string      `json:"recipient_account_type" validate:"required"`
	RecipientBankName         string      `json:"recipient_bank_name" validate:"required"`
	RecipientBankCode         string      `json:"recipient_bank_code" validate:"required"`
	RecipientBankAccountNo    string      `json:"recipient_bank_account_no" validate:"required"`
	RecipientBankAccountOwner string      `json:"recipient_bank_account_owner" validate:"required"`
	RecipientNotes            *string     `json:"recipient_notes,omitempty" validate:"omitempty"`
	Rate                      json.Number `json:"rate" validate:"required"`
	FeePercent                json.Number `json:"fee_percent" validate:"required"`
	FeeFixed                  json.Number `json:"fee_fixed" validate:"required"`
	PaymentMethod             string      `json:"payment_method" validate:"required"`
	TransferType              string      `json:"transfer_type,omitempty" validate:"omitempty"`
	TransferReference         *string     `json:"transfer_reference,omitempty" validate:"omitempty"`
	FundSource                *string     `json:"fund_source,omitempty" validate:"omitempty"`
	Purpose                   *string     `json:"purpose,omitempty" validate:"omitempty"`
	Notes                     *string     `json:"notes,omitempty" validate:"omitempty"`
	Files                     *[]string   `json:"files,omitempty" validate:"omitempty"`
}

type GetRemittanceFilter struct {
	Status       *string `json:"status,omitempty" query:"status"`
	Type         *string `json:"type,omitempty" query:"type"`
	Search       *string `json:"search,omitempty" query:"search"`
	Start        *string `json:"start,omitempty" query:"start"`
	End          *string `json:"end,omitempty" query:"end"`
	FromCurrency *string `json:"from_currency,omitempty" query:"from_currency"`
	ToCurrency   *string `json:"to_currency,omitempty" query:"to_currency"`
}
