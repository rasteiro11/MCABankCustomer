package repository

import (
	"context"

	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository/models"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository/models/mappers"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

var _ CustomerRepository = (*customerRepository)(nil)

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var ms []models.Customer
	if err := r.db.WithContext(ctx).Find(&ms).Error; err != nil {
		return nil, err
	}
	customers := make([]domain.Customer, 0, len(ms))
	for _, m := range ms {
		customers = append(customers, *mappers.ToDomain(&m))
	}
	return customers, nil
}

func (r *customerRepository) FindByID(ctx context.Context, id uint) (*domain.Customer, error) {
	var m models.Customer
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return mappers.ToDomain(&m), nil
}

func (r *customerRepository) Create(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	m := mappers.FromDomain(customer)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}
	return mappers.ToDomain(m), nil
}

func (r *customerRepository) CreateWithCallback(
	ctx context.Context,
	customer *domain.Customer,
	fn func(*domain.Customer) error,
) (*domain.Customer, error) {
	var created *domain.Customer

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		m := mappers.FromDomain(customer)

		if err := tx.Create(m).Error; err != nil {
			return err
		}

		created = mappers.ToDomain(m)

		if err := fn(created); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r *customerRepository) Update(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.Customer{}).
		Where("id = ?", customer.ID).
		Updates(models.Customer{
			Nome:  customer.Nome,
			Email: customer.Email,
		}).Error; err != nil {
		return nil, err
	}

	return r.FindByID(ctx, customer.ID)
}

func (r *customerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Customer{}, id).Error
}
