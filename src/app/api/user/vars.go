package user

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
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,eqField=NewPassword"`
}

type CreateContactRequest struct {
	Name          string  `json:"name" validate:"required"`
	Birthdate     *string `json:"birthdate,omitempty" validate:"omitempty"`
	Gender        *string `json:"gender,omitempty" validate:"omitempty"`
	Country       *string `json:"country,omitempty" validate:"omitempty"`
	Region        *string `json:"region,omitempty" validate:"omitempty"`
	City          *string `json:"city,omitempty" validate:"omitempty"`
	Address       *string `json:"address,omitempty" validate:"omitempty"`
	Email         *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode     *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo       *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	IdentityType  *string `json:"identity_type,omitempty" validate:"omitempty"`
	IdentityNo    *string `json:"identity_no,omitempty" validate:"omitempty"`
	Occupation    *string `json:"occupation,omitempty" validate:"omitempty"`
	IdentityImage *string `json:"identity_image,omitempty" validate:"omitempty"`
	PortraitImage *string `json:"portrait_image,omitempty" validate:"omitempty"`
	Notes         *string `json:"notes,omitempty" validate:"omitempty"`
}

type UpdateContactRequest struct {
	Name          string  `json:"name" validate:"required"`
	Birthdate     *string `json:"birthdate,omitempty" validate:"omitempty"`
	Gender        *string `json:"gender,omitempty" validate:"omitempty"`
	Country       *string `json:"country,omitempty" validate:"omitempty"`
	Region        *string `json:"region,omitempty" validate:"omitempty"`
	City          *string `json:"city,omitempty" validate:"omitempty"`
	Address       *string `json:"address,omitempty" validate:"omitempty"`
	Email         *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode     *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo       *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	IdentityType  *string `json:"identity_type,omitempty" validate:"omitempty"`
	IdentityNo    *string `json:"identity_no,omitempty" validate:"omitempty"`
	Occupation    *string `json:"occupation,omitempty" validate:"omitempty"`
	IdentityImage *string `json:"identity_image,omitempty" validate:"omitempty"`
	PortraitImage *string `json:"portrait_image,omitempty" validate:"omitempty"`
	Notes         *string `json:"notes,omitempty" validate:"omitempty"`
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
	FromContact  string  `json:"from_contact" validate:"required"`
	FromCurrency string  `json:"from_currency" validate:"required"`
	FromAmount   float64 `json:"from_amount" validate:"required"`
	ToContact    string  `json:"to_contact" validate:"required"`
	ToCurrency   string  `json:"to_currency" validate:"required"`
	ToAmount     float64 `json:"to_amount" validate:"required"`
	ToAccount    string  `json:"to_account" validate:"required"`
	RateCurrency string  `json:"rate_currency" validate:"required"`
	Rate         float64 `json:"rate" validate:"required"`
	FeeCurrency  string  `json:"fee_currency" validate:"required"`
	FeeAmount    float64 `json:"fee_amount" validate:"required"`
	TransferType string  `json:"transfer_type" validate:"required"`
	TransferRef  *string `json:"transfer_ref,omitempty" validate:"omitempty"`
	TransferDate *string `json:"transfer_date,omitempty" validate:"omitempty"`
	FundSource   *string `json:"fund_source,omitempty" validate:"omitempty"`
	Purpose      *string `json:"purpose,omitempty" validate:"omitempty"`
	Notes        *string `json:"notes,omitempty" validate:"omitempty"`
}
