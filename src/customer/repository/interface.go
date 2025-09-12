package repository

import (
	"context"

	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
)

type (
	CustomerRepository interface {
		FindAll(ctx context.Context) ([]domain.Customer, error)
		FindByID(ctx context.Context, id uint) (*domain.Customer, error)
		Create(ctx context.Context, customer *domain.Customer) (*domain.Customer, error)
		CreateWithCallback(ctx context.Context, customer *domain.Customer, fn func(*domain.Customer) error) (*domain.Customer, error)
		Update(ctx context.Context, customer *domain.Customer) (*domain.Customer, error)
		Delete(ctx context.Context, id uint) error
	}
)
