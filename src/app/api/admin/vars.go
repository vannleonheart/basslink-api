package admin

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

type CreateUserRequest struct {
	Role                 string  `json:"role" validate:"required"`
	Username             string  `json:"username" validate:"required,min=5,max=50"`
	Name                 string  `json:"name" validate:"required"`
	Email                *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode            *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo              *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Password             string  `json:"password" validate:"required,min=8"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type UpdateUserRequest struct {
	Role                 string  `json:"role" validate:"required"`
	Username             string  `json:"username" validate:"required,min=5,max=50"`
	Name                 string  `json:"name" validate:"required"`
	Email                *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneCode            *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo              *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Password             *string `json:"password,omitempty" validate:"omitempty,min=8"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty" validate:"omitempty,eqfield=Password"`
}

type CreateAgentRequest struct {
	AgentName            string  `json:"agent_name" validate:"required,min=5,max=100"`
	Country              *string `json:"country,omitempty" validate:"omitempty"`
	Region               *string `json:"region,omitempty" validate:"omitempty"`
	City                 *string `json:"city,omitempty" validate:"omitempty"`
	Address              *string `json:"address,omitempty" validate:"omitempty"`
	PhoneCode            *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo              *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Email                *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	Website              *string `json:"website,omitempty" validate:"omitempty,url"`
	Name                 string  `json:"name" validate:"required,min=5,max=100"`
	Username             string  `json:"username" validate:"required,min=5,max=50"`
	Password             string  `json:"password" validate:"required,min=8"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type UpdateAgentRequest struct {
	AgentName string  `json:"agent_name" validate:"required,min=5,max=100"`
	Country   *string `json:"country,omitempty" validate:"omitempty"`
	Region    *string `json:"region,omitempty" validate:"omitempty"`
	City      *string `json:"city,omitempty" validate:"omitempty"`
	Address   *string `json:"address,omitempty" validate:"omitempty"`
	PhoneCode *string `json:"phone_code,omitempty" validate:"omitempty,max=5"`
	PhoneNo   *string `json:"phone_no,omitempty" validate:"omitempty,max=15"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email,max=100"`
	Website   *string `json:"website,omitempty" validate:"omitempty,url"`
}

type GetDisbursementFilter struct {
	Status *string `json:"status,omitempty" query:"status"`
	Search *string `json:"search,omitempty" query:"search"`
	Start  *string `json:"start,omitempty" query:"start"`
	End    *string `json:"end,omitempty" query:"end"`
}
