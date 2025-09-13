package http

type createCustomerRequest struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type updateCustomerRequest struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type customerResponse struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
