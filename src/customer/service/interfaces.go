package service

import (
	"context"

	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
)

type CustomerService interface {
	GetAll(ctx context.Context) ([]domain.Customer, error)
	GetByID(ctx context.Context, id uint) (*domain.Customer, error)
	Create(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Delete(ctx context.Context, id uint) error
}
