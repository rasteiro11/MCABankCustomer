package mappers

import (
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/repository/models"
	"gorm.io/gorm"
)

func FromDomain(c *domain.Customer) *models.Customer {
	if c == nil {
		return nil
	}
	return &models.Customer{
		Model: gorm.Model{ID: c.ID},
		Nome:  c.Nome,
		Email: c.Email,
	}
}

func ToDomain(m *models.Customer) *domain.Customer {
	if m == nil {
		return nil
	}
	return &domain.Customer{
		ID:    m.ID,
		Nome:  m.Nome,
		Email: m.Email,
	}
}
