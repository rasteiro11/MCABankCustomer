package service

import (
	"context"

	pbPaymentClient "github.com/rasteiro11/MCABankCustomer/gen/proto/go/payment"
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository"
	"github.com/rasteiro11/PogCore/pkg/logger"
)

type customerService struct {
	repo          repository.CustomerRepository
	paymentClient pbPaymentClient.BalanceServiceClient
}

func NewCustomerService(repo repository.CustomerRepository, paymentClient pbPaymentClient.BalanceServiceClient) CustomerService {
	return &customerService{repo: repo, paymentClient: paymentClient}
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
	m, err := s.repo.CreateWithCallback(ctx, (c), func(c *domain.Customer) error {
		logger.Of(ctx).Infof("Creating balance for customer ID %d", c.ID)
		if _, err := s.paymentClient.CreateBalance(ctx, &pbPaymentClient.CreateBalanceRequest{
			CustomerId: uint32(c.ID),
		}); err != nil {
			return err
		}
		logger.Of(ctx).Infof("Balance created successfully for customer ID %d", c.ID)

		return nil
	})
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
