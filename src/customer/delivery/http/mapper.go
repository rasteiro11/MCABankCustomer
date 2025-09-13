package http

import (
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
)

func MapCreateRequestToDomain(req *createCustomerRequest) *domain.Customer {
	return &domain.Customer{
		Nome:  req.Nome,
		Email: req.Email,
	}
}

func MapCustomerToHTTP(c *domain.Customer) *customerResponse {
	if c == nil {
		return nil
	}
	return &customerResponse{
		ID:    c.ID,
		Nome:  c.Nome,
		Email: c.Email,
	}
}

func MapCustomersToHTTP(customers []domain.Customer) []*customerResponse {
	out := make([]*customerResponse, 0, len(customers))
	for _, c := range customers {
		out = append(out, MapCustomerToHTTP(&c))
	}
	return out
}
