package http

import (
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
)

func MapCreateRequestToDomain(req *CreateCustomerRequest) *domain.Customer {
	return &domain.Customer{
		Nome:  req.Nome,
		Email: req.Email,
	}
}

func MapCustomerToHTTP(c *domain.Customer) *CustomerResponse {
	if c == nil {
		return nil
	}
	return &CustomerResponse{
		ID:    c.ID,
		Nome:  c.Nome,
		Email: c.Email,
	}
}

func MapCustomersToHTTP(customers []domain.Customer) []*CustomerResponse {
	out := make([]*CustomerResponse, 0, len(customers))
	for _, c := range customers {
		out = append(out, MapCustomerToHTTP(&c))
	}
	return out
}
