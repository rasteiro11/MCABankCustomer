package http

type CreateCustomerRequest struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateCustomerRequest struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CustomerResponse struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
