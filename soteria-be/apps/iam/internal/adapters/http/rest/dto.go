package rest

type ApiResponseDto struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ApiErrorResponseDto struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type CreatePermissionDto struct {
	Resource string `json:"resource" validate:"required"`
	Action   string `json:"action" validate:"required"`
}

type UpdatePermissionDto struct {
	Resource string `json:"resource" validate:"required"`
	Action   string `json:"action" validate:"required"`
}

type CreateUserDto struct {
	Name            string              `json:"name" validate:"required"`
	Email           string              `json:"email" validate:"required,email"`
	PhoneNumber     string              `json:"phoneNumber" validate:"required,e164"`
	Password        string              `json:"password" validate:"required,min=6"`
	TimeZone        string              `json:"timeZone" validate:"omitempty"`
	DeliveryMethods []string            `json:"deliveryMethods" validate:"omitempty,dive,eq=email|eq=sms|eq=call"`
	Permissions     []UserPermissionDto `json:"permissions" validate:"omitempty,dive"`
}

type UpdateUserDto struct {
	Name            string              `json:"name" validate:"required"`
	Email           string              `json:"email" validate:"required,email"`
	PhoneNumber     string              `json:"phoneNumber" validate:"required,e164"`
	TimeZone        string              `json:"timeZone" validate:"omitempty"`
	DeliveryMethods []string            `json:"deliveryMethods" validate:"omitempty,dive,eq=email|eq=sms|eq=call"`
	Permissions     []UserPermissionDto `json:"permissions" validate:"omitempty,dive"`
}

type UserPermissionDto struct {
	Resource string `json:"resource" validate:"required"`
	Action   string `json:"action" validate:"required"`
}

type CreateApiKeyDto struct {
	Name        string              `json:"name" validate:"required"`
	Permissions []UserPermissionDto `json:"permissions" validate:"required,min=1,dive"`
}

type VerifyApiKeyDto struct {
	Key string `json:"key" validate:"required"`
}

type SignUpDto struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=6"`
}

type SignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenDto struct {
	Token string `json:"token" validate:"required"`
}
