package service

import (
	"context"

	pbPaymentClient "github.com/rasteiro11/MCABankCustomer/gen/proto/go/payment"
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("customer-service")

type customerService struct {
	repo          repository.CustomerRepository
	paymentClient pbPaymentClient.BalanceServiceClient
}

func NewCustomerService(repo repository.CustomerRepository, paymentClient pbPaymentClient.BalanceServiceClient) CustomerService {
	return &customerService{repo: repo, paymentClient: paymentClient}
}

func (s *customerService) GetAll(ctx context.Context) ([]domain.Customer, error) {
	ctx, span := tracer.Start(ctx, "GetAll")
	defer span.End()

	customers, err := s.repo.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(attribute.Int("customer.count", len(customers)))
	return customers, nil
}

func (s *customerService) GetByID(ctx context.Context, id uint) (*domain.Customer, error) {
	ctx, span := tracer.Start(ctx, "GetByID")
	defer span.End()
	span.SetAttributes(attribute.Int("customer.id", int(id)))

	m, err := s.repo.FindByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return m, nil
}

func (s *customerService) Create(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	ctx, span := tracer.Start(ctx, "CreateCustomer")
	defer span.End()

	m, err := s.repo.CreateWithCallback(ctx, c, func(c *domain.Customer) error {
		_, subSpan := tracer.Start(ctx, "CreateBalance")
		defer subSpan.End()

		logger.Of(ctx).Infof("Creating balance for customer ID %d", c.ID)
		if _, err := s.paymentClient.CreateBalance(ctx, &pbPaymentClient.CreateBalanceRequest{
			CustomerId: uint32(c.ID),
		}); err != nil {
			subSpan.RecordError(err)
			return err
		}
		logger.Of(ctx).Infof("Balance created successfully for customer ID %d", c.ID)
		return nil
	})
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return m, nil
}

func (s *customerService) Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	ctx, span := tracer.Start(ctx, "UpdateCustomer")
	defer span.End()
	span.SetAttributes(attribute.Int("customer.id", int(c.ID)))

	m, err := s.repo.Update(ctx, c)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return m, nil
}

func (s *customerService) Delete(ctx context.Context, id uint) error {
	ctx, span := tracer.Start(ctx, "DeleteCustomer")
	defer span.End()
	span.SetAttributes(attribute.Int("customer.id", int(id)))

	if err := s.repo.Delete(ctx, id); err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}
