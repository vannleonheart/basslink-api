package common

import "encoding/json"

type GetRateRequest struct {
	FromCurrency string  `json:"from_currency" validate:"required"`
	ToCurrency   string  `json:"to_currency" validate:"required"`
	FromAmount   *string `json:"from_amount,omitempty" validate:"omitempty"`
	ToAmount     *string `json:"to_amount,omitempty" validate:"omitempty"`
}

type GetRequestResponse struct {
	FromCurrency string      `json:"from_currency"`
	ToCurrency   string      `json:"to_currency"`
	FromAmount   json.Number `json:"from_amount"`
	ToAmount     json.Number `json:"to_amount"`
	Rate         json.Number `json:"rate"`
	FeePercent   json.Number `json:"fee_percent"`
	FeeFixed     json.Number `json:"fee_fixed"`
	TotalFee     json.Number `json:"total_fee"`
}

type CreateAppointmentRequest struct {
	Name    string  `json:"name" validate:"required,min=5,max=100"`
	Company *string `json:"company,omitempty" validate:"omitempty"`
	Email   string  `json:"email" validate:"required,email,max=100"`
	Phone   string  `json:"phone" validate:"required,min=5,max=15"`
	Service string  `json:"service" validate:"required"`
	Date    string  `json:"date" validate:"required"`
	Time    string  `json:"time" validate:"required"`
	Notes   *string `json:"notes,omitempty" validate:"omitempty"`
	Token   string  `json:"token" validate:"required"`
}

type TransactionSearchRequest struct {
	TransactionId string `json:"transaction_id" validate:"required"`
	SenderName    string `json:"sender_name" validate:"required"`
	RecipientName string `json:"recipient_name" validate:"required"`
}

type CreateRemittanceRequest struct {
	FromCurrency              string      `json:"from_currency" validate:"required"`
	FromAmount                json.Number `json:"from_amount" validate:"required"`
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
	ToCurrency                string      `json:"to_currency,omitempty" validate:"required"`
	ToAmount                  json.Number `json:"to_amount" validate:"required"`
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
	RecipientAccountType      string      `json:"recipient_account_type" validate:"required,oneof=bank_account ewallet"`
	RecipientBankName         string      `json:"recipient_bank_name" validate:"required"`
	RecipientBankCode         *string     `json:"recipient_bank_code,omitempty" validate:"omitempty"`
	RecipientBankAccountNo    string      `json:"recipient_bank_account_no" validate:"required"`
	RecipientBankAccountOwner string      `json:"recipient_bank_account_owner" validate:"required"`
	RecipientNotes            *string     `json:"recipient_notes,omitempty" validate:"omitempty"`
	TransferType              string      `json:"transfer_type,omitempty" validate:"omitempty"`
	TransferReference         *string     `json:"transfer_reference,omitempty" validate:"omitempty"`
	FundSource                *string     `json:"fund_source,omitempty" validate:"omitempty"`
	Purpose                   *string     `json:"purpose,omitempty" validate:"omitempty"`
	Notes                     *string     `json:"notes,omitempty" validate:"omitempty"`
	Files                     *[]string   `json:"files,omitempty" validate:"omitempty"`
	Token                     string      `json:"token" validate:"required"`
}

type BankInfo struct {
	BankName     string
	BankCode     string
	SwiftCode    string
	AccountNo    string
	AccountOwner string
	Currency     string
}
