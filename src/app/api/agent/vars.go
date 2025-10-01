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

type CreateUserRequest struct {
	CustomerType         string  `json:"customer_type" validate:"required,oneof=individual institution"`
	CustomerName         string  `json:"customer_name" validate:"required,min=5,max=100"`
	CustomerGender       *string `json:"customer_gender,omitempty" validate:"omitempty,oneof=male female"`
	CustomerBirthdate    *string `json:"customer_birthdate,omitempty" validate:"omitempty"`
	CustomerCitizenship  string  `json:"customer_citizenship" validate:"required"`
	CustomerCountry      *string `json:"customer_country,omitempty" validate:"omitempty"`
	CustomerRegion       *string `json:"customer_region,omitempty" validate:"omitempty"`
	CustomerCity         *string `json:"customer_city,omitempty" validate:"omitempty"`
	CustomerAddress      *string `json:"customer_address,omitempty" validate:"omitempty"`
	CustomerEmail        *string `json:"customer_email,omitempty" validate:"omitempty,email,max=100"`
	CustomerPhoneCode    *string `json:"customer_phone_code,omitempty" validate:"omitempty,max=5"`
	CustomerPhoneNo      *string `json:"customer_phone_no,omitempty" validate:"omitempty,max=15"`
	CustomerIdentityType string  `json:"customer_identity_type,omitempty" validate:"required"`
	CustomerIdentityNo   string  `json:"customer_identity_no,omitempty" validate:"required,min=3,max=100"`
	CustomerOccupation   *string `json:"customer_occupation,omitempty" validate:"omitempty"`
	CustomerNotes        *string `json:"customer_notes,omitempty" validate:"omitempty"`
	CustomerDocuments    *[]struct {
		DocumentData string  `json:"document_data" validate:"required"`
		DocumentType string  `json:"document_type" validate:"required"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"customer_documents" validate:"omitempty,dive"`
	Username             *string `json:"username,omitempty" validate:"omitempty,min=5,max=50"`
	Password             *string `json:"password,omitempty" validate:"omitempty,min=8"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty" validate:"omitempty,eqfield=Password"`
}

type UpdateUserRequest struct {
	CustomerType         string  `json:"customer_type" validate:"required,oneof=individual institution"`
	CustomerName         string  `json:"customer_name" validate:"required,min=5,max=100"`
	CustomerGender       *string `json:"customer_gender,omitempty" validate:"omitempty,oneof=male female"`
	CustomerBirthdate    *string `json:"customer_birthdate,omitempty" validate:"omitempty"`
	CustomerCitizenship  string  `json:"customer_citizenship" validate:"required"`
	CustomerCountry      *string `json:"customer_country,omitempty" validate:"omitempty"`
	CustomerRegion       *string `json:"customer_region,omitempty" validate:"omitempty"`
	CustomerCity         *string `json:"customer_city,omitempty" validate:"omitempty"`
	CustomerAddress      *string `json:"customer_address,omitempty" validate:"omitempty"`
	CustomerEmail        *string `json:"customer_email,omitempty" validate:"omitempty,email,max=100"`
	CustomerPhoneCode    *string `json:"customer_phone_code,omitempty" validate:"omitempty,max=5"`
	CustomerPhoneNo      *string `json:"customer_phone_no,omitempty" validate:"omitempty,max=15"`
	CustomerIdentityType string  `json:"customer_identity_type,omitempty" validate:"required,oneof=passport national_id other"`
	CustomerIdentityNo   string  `json:"customer_identity_no,omitempty" validate:"required,min=3,max=100"`
	CustomerOccupation   *string `json:"customer_occupation,omitempty" validate:"omitempty"`
	CustomerNotes        *string `json:"customer_notes,omitempty" validate:"omitempty"`
	CustomerDocuments    *[]struct {
		Id           *string `json:"id,omitempty" validate:"omitempty"`
		DocumentData string  `json:"document_data" validate:"required"`
		DocumentType string  `json:"document_type" validate:"required"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"customer_documents" validate:"omitempty,dive"`
	Username             *string `json:"username,omitempty" validate:"omitempty,min=5,max=50"`
	Password             *string `json:"password,omitempty" validate:"omitempty,min=8"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty" validate:"omitempty,eqfield=Password"`
}

type CreateContactRequest struct {
	ContactType         string  `json:"contact_type" validate:"required,oneof=individual institution"`
	ContactName         string  `json:"contact_name" validate:"required,min=5,max=100"`
	ContactGender       *string `json:"contact_gender,omitempty" validate:"omitempty,oneof=male female"`
	ContactBirthdate    *string `json:"contact_birthdate,omitempty" validate:"omitempty"`
	ContactCitizenship  string  `json:"contact_citizenship" validate:"required"`
	ContactCountry      *string `json:"contact_country,omitempty" validate:"omitempty"`
	ContactRegion       *string `json:"contact_region,omitempty" validate:"omitempty"`
	ContactCity         *string `json:"contact_city,omitempty" validate:"omitempty"`
	ContactAddress      *string `json:"contact_address,omitempty" validate:"omitempty"`
	ContactEmail        *string `json:"contact_email,omitempty" validate:"omitempty,email,max=100"`
	ContactPhoneCode    *string `json:"contact_phone_code,omitempty" validate:"omitempty,max=5"`
	ContactPhoneNo      *string `json:"contact_phone_no,omitempty" validate:"omitempty,max=15"`
	ContactIdentityType string  `json:"contact_identity_type,omitempty" validate:"required,oneof=passport national_id other"`
	ContactIdentityNo   string  `json:"contact_identity_no,omitempty" validate:"required,min=3,max=100"`
	ContactOccupation   *string `json:"contact_occupation,omitempty" validate:"omitempty"`
	ContactNotes        *string `json:"contact_notes,omitempty" validate:"omitempty"`
	ContactDocuments    *[]struct {
		DocumentData string  `json:"document_data" validate:"required"`
		DocumentType string  `json:"document_type" validate:"required"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"contact_documents" validate:"omitempty,dive"`
	ContactAccounts *[]struct {
		BankName        string  `json:"bank_name" validate:"required"`
		BankAccountNo   string  `json:"bank_account_no" validate:"required"`
		BankAccountName string  `json:"bank_account_name" validate:"required"`
		BankCountry     string  `json:"bank_country" validate:"required"`
		BankCode        *string `json:"bank_code,omitempty" validate:"omitempty"`
		BankSwiftCode   *string `json:"bank_swift_code,omitempty" validate:"omitempty"`
		BankAddress     *string `json:"bank_address,omitempty" validate:"omitempty"`
		BankEmail       *string `json:"bank_email,omitempty" validate:"omitempty,email,max=100"`
		BankPhoneCode   *string `json:"bank_phone_code,omitempty" validate:"omitempty,max=5"`
		BankPhoneNo     *string `json:"bank_phone_no,omitempty" validate:"omitempty,max=15"`
		BankWebsite     *string `json:"bank_website,omitempty" validate:"omitempty,url"`
		BankNotes       *string `json:"bank_notes,omitempty" validate:"omitempty"`
	} `json:"contact_accounts" validate:"omitempty,dive"`
}

type UpdateContactRequest struct {
	ContactType         string  `json:"contact_type" validate:"required,oneof=individual institution"`
	ContactName         string  `json:"contact_name" validate:"required,min=5,max=100"`
	ContactGender       *string `json:"contact_gender,omitempty" validate:"omitempty,oneof=male female"`
	ContactBirthdate    *string `json:"contact_birthdate,omitempty" validate:"omitempty"`
	ContactCitizenship  string  `json:"contact_citizenship" validate:"required"`
	ContactCountry      *string `json:"contact_country,omitempty" validate:"omitempty"`
	ContactRegion       *string `json:"contact_region,omitempty" validate:"omitempty"`
	ContactCity         *string `json:"contact_city,omitempty" validate:"omitempty"`
	ContactAddress      *string `json:"contact_address,omitempty" validate:"omitempty"`
	ContactEmail        *string `json:"contact_email,omitempty" validate:"omitempty,email,max=100"`
	ContactPhoneCode    *string `json:"contact_phone_code,omitempty" validate:"omitempty,max=5"`
	ContactPhoneNo      *string `json:"contact_phone_no,omitempty" validate:"omitempty,max=15"`
	ContactIdentityType string  `json:"contact_identity_type,omitempty" validate:"required,oneof=passport national_id other"`
	ContactIdentityNo   string  `json:"contact_identity_no,omitempty" validate:"required,min=3,max=100"`
	ContactOccupation   *string `json:"contact_occupation,omitempty" validate:"omitempty"`
	ContactNotes        *string `json:"contact_notes,omitempty" validate:"omitempty"`
	ContactDocuments    *[]struct {
		Id           *string `json:"id,omitempty" validate:"omitempty"`
		DocumentData string  `json:"document_data" validate:"required"`
		DocumentType string  `json:"document_type" validate:"required"`
		Notes        *string `json:"notes,omitempty" validate:"omitempty"`
		IsVerified   *bool   `json:"is_verified" validate:"omitempty"`
	} `json:"contact_documents" validate:"omitempty,dive"`
	ContactAccounts *[]struct {
		Id              *string `json:"id,omitempty" validate:"omitempty"`
		BankName        string  `json:"bank_name" validate:"required"`
		BankAccountNo   string  `json:"bank_account_no" validate:"required"`
		BankAccountName string  `json:"bank_account_name" validate:"required"`
		BankCountry     string  `json:"bank_country" validate:"required"`
		BankCode        *string `json:"bank_code,omitempty" validate:"omitempty"`
		BankSwiftCode   *string `json:"bank_swift_code,omitempty" validate:"omitempty"`
		BankAddress     *string `json:"bank_address,omitempty" validate:"omitempty"`
		BankEmail       *string `json:"bank_email,omitempty" validate:"omitempty,email,max=100"`
		BankPhoneCode   *string `json:"bank_phone_code,omitempty" validate:"omitempty,max=5"`
		BankPhoneNo     *string `json:"bank_phone_no,omitempty" validate:"omitempty,max=15"`
		BankWebsite     *string `json:"bank_website,omitempty" validate:"omitempty,url"`
		BankNotes       *string `json:"bank_notes,omitempty" validate:"omitempty"`
	} `json:"contact_accounts" validate:"omitempty,dive"`
}

type CreateContactDocumentRequest struct {
	DocumentType string  `json:"document_type" validate:"required"`
	DocumentData string  `json:"document_data" validate:"required"`
	Notes        *string `json:"notes,omitempty" validate:"omitempty"`
}

type UpdateContactDocumentRequest struct {
	DocumentType string  `json:"document_type" validate:"required"`
	DocumentData string  `json:"document_data" validate:"required"`
	Notes        *string `json:"notes,omitempty" validate:"omitempty"`
}

type CreateContactAccountRequest struct {
	AccountType string  `json:"account_type" validate:"required"`
	BankId      string  `json:"bank_id" validate:"required"`
	BankName    string  `json:"bank_name" validate:"required"`
	BankCode    *string `json:"bank_code,omitempty" validate:"omitempty"`
	BankSwift   *string `json:"bank_swift,omitempty" validate:"omitempty"`
	OwnerName   string  `json:"owner_name" validate:"required"`
	No          string  `json:"no" validate:"required"`
	Country     *string `json:"country,omitempty" validate:"omitempty"`
	Region      *string `json:"region,omitempty" validate:"omitempty"`
	City        *string `json:"city,omitempty" validate:"omitempty"`
	Address     *string `json:"address,omitempty" validate:"omitempty"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	Website     *string `json:"website,omitempty" validate:"omitempty,url"`
	PhoneCode   *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo     *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Notes       *string `json:"notes,omitempty" validate:"omitempty"`
}

type UpdateContactAccountRequest struct {
	AccountType string  `json:"account_type" validate:"required"`
	BankId      string  `json:"bank_id" validate:"required"`
	BankName    string  `json:"bank_name" validate:"required"`
	BankCode    *string `json:"bank_code,omitempty" validate:"omitempty"`
	BankSwift   *string `json:"bank_swift,omitempty" validate:"omitempty"`
	OwnerName   string  `json:"owner_name" validate:"required"`
	No          string  `json:"no" validate:"required"`
	Country     *string `json:"country,omitempty" validate:"omitempty"`
	Region      *string `json:"region,omitempty" validate:"omitempty"`
	City        *string `json:"city,omitempty" validate:"omitempty"`
	Address     *string `json:"address,omitempty" validate:"omitempty"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	Website     *string `json:"website,omitempty" validate:"omitempty,url"`
	PhoneCode   *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo     *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Notes       *string `json:"notes,omitempty" validate:"omitempty"`
}

type CreateDisbursementRequest struct {
	FromCurrency            string      `json:"from_currency" validate:"required"`
	FromAmount              json.Number `json:"from_amount" validate:"required"`
	FromCustomer            *string     `json:"from_customer,omitempty" validate:"omitempty"`
	CustomerType            string      `json:"customer_type" validate:"required,oneof=individual institution"`
	CustomerName            string      `json:"customer_name" validate:"required,min=5,max=100"`
	CustomerGender          *string     `json:"customer_gender,omitempty" validate:"omitempty,oneof=male female"`
	CustomerBirthdate       *string     `json:"customer_birthdate,omitempty" validate:"omitempty"`
	CustomerCitizenship     string      `json:"customer_citizenship" validate:"required"`
	CustomerIdentityType    string      `json:"customer_identity_type" validate:"required"`
	CustomerIdentityNo      string      `json:"customer_identity_no" validate:"required,min=3,max=100"`
	CustomerOccupation      *string     `json:"customer_occupation,omitempty" validate:"omitempty"`
	CustomerCountry         *string     `json:"customer_country,omitempty" validate:"omitempty"`
	CustomerRegion          *string     `json:"customer_region,omitempty" validate:"omitempty"`
	CustomerCity            *string     `json:"customer_city,omitempty" validate:"omitempty"`
	CustomerAddress         *string     `json:"customer_address,omitempty" validate:"omitempty"`
	CustomerEmail           *string     `json:"customer_email,omitempty" validate:"omitempty,email,max=100"`
	CustomerPhoneCode       *string     `json:"customer_phone_code,omitempty" validate:"omitempty,max=5"`
	CustomerPhoneNo         *string     `json:"customer_phone_no,omitempty" validate:"omitempty,max=15"`
	CustomerNotes           *string     `json:"customer_notes,omitempty" validate:"omitempty"`
	CustomerUpdate          *bool       `json:"customer_update,omitempty" validate:"omitempty"`
	ToCurrency              string      `json:"to_currency,omitempty" validate:"required"`
	ToAmount                json.Number `json:"to_amount" validate:"required"`
	ToContact               *string     `json:"to_contact,omitempty" validate:"omitempty"`
	BeneficiaryType         string      `json:"beneficiary_type" validate:"required,oneof=individual institution"`
	BeneficiaryName         string      `json:"beneficiary_name" validate:"required,min=5,max=100"`
	BeneficiaryGender       *string     `json:"beneficiary_gender,omitempty" validate:"omitempty,oneof=male female"`
	BeneficiaryBirthdate    *string     `json:"beneficiary_birthdate,omitempty" validate:"omitempty"`
	BeneficiaryCitizenship  string      `json:"beneficiary_citizenship" validate:"required"`
	BeneficiaryIdentityType string      `json:"beneficiary_identity_type" validate:"required"`
	BeneficiaryIdentityNo   string      `json:"beneficiary_identity_no" validate:"required,min=3,max=100"`
	BeneficiaryOccupation   *string     `json:"beneficiary_occupation,omitempty" validate:"omitempty"`
	BeneficiaryCountry      *string     `json:"beneficiary_country,omitempty" validate:"omitempty"`
	BeneficiaryRegion       *string     `json:"beneficiary_region,omitempty" validate:"omitempty"`
	BeneficiaryCity         *string     `json:"beneficiary_city,omitempty" validate:"omitempty"`
	BeneficiaryAddress      *string     `json:"beneficiary_address,omitempty" validate:"omitempty"`
	BeneficiaryEmail        *string     `json:"beneficiary_email,omitempty" validate:"omitempty,email,max=100"`
	BeneficiaryPhoneCode    *string     `json:"beneficiary_phone_code,omitempty" validate:"omitempty,max=5"`
	BeneficiaryPhoneNo      *string     `json:"beneficiary_phone_no,omitempty" validate:"omitempty,max=15"`
	BeneficiaryNotes        *string     `json:"beneficiary_notes,omitempty" validate:"omitempty"`
	BeneficiaryUpdate       *bool       `json:"beneficiary_update,omitempty" validate:"omitempty"`
	BeneficiaryRelationship *string     `json:"beneficiary_relationship,omitempty" validate:"omitempty"`
	ToAccount               *string     `json:"to_account,omitempty" validate:"omitempty"`
	BankName                string      `json:"bank_name" validate:"required"`
	BankAccountNo           string      `json:"bank_account_no" validate:"required"`
	BankAccountName         string      `json:"bank_account_name" validate:"required"`
	BankCountry             string      `json:"bank_country" validate:"required"`
	BankCode                *string     `json:"bank_code,omitempty" validate:"omitempty"`
	BankSwiftCode           *string     `json:"bank_swift_code,omitempty" validate:"omitempty"`
	BankAddress             *string     `json:"bank_address,omitempty" validate:"omitempty"`
	BankEmail               *string     `json:"bank_email,omitempty" validate:"omitempty,email,max=100"`
	BankPhoneCode           *string     `json:"bank_phone_code,omitempty" validate:"omitempty,max=5"`
	BankPhoneNo             *string     `json:"bank_phone_no,omitempty" validate:"omitempty,max=15"`
	BankWebsite             *string     `json:"bank_website,omitempty" validate:"omitempty,url"`
	BankNotes               *string     `json:"bank_notes,omitempty" validate:"omitempty"`
	BankUpdate              *bool       `json:"bank_update,omitempty" validate:"omitempty"`
	Rate                    json.Number `json:"rate" validate:"required"`
	FeePercent              json.Number `json:"fee_percent" validate:"required"`
	FeeFixed                json.Number `json:"fee_fixed" validate:"required"`
	TransferType            *string     `json:"transfer_type,omitempty" validate:"omitempty"`
	TransferDate            *string     `json:"transfer_date,omitempty" validate:"omitempty"`
	TransferReference       *string     `json:"transfer_reference,omitempty" validate:"omitempty"`
	FundSource              *string     `json:"fund_source,omitempty" validate:"omitempty"`
	Purpose                 *string     `json:"purpose,omitempty" validate:"omitempty"`
	Notes                   *string     `json:"notes,omitempty" validate:"omitempty"`
	Files                   *[]string   `json:"files,omitempty" validate:"omitempty"`
}

type GetDisbursementFilter struct {
	Status *string `json:"status,omitempty" query:"status"`
	Search *string `json:"search,omitempty" query:"search"`
	Start  *string `json:"start,omitempty" query:"start"`
	End    *string `json:"end,omitempty" query:"end"`
}
