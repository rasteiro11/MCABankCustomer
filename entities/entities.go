package entities

import "github.com/rasteiro11/MCABankCustomer/src/customer/repository/models"

func GetEntities() []any {
	return []any{
		&models.Customer{},
	}
}
