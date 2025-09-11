package service

import (
	"context"

	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository"
)

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetAll(ctx context.Context) ([]domain.Customer, error) {
	customers, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *customerService) GetByID(ctx context.Context, id uint) (*domain.Customer, error) {
	m, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *customerService) Create(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	m, err := s.repo.Create(ctx, c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *customerService) Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	m, err := s.repo.Update(ctx, c)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *customerService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
