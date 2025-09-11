package domain

type Customer struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome  string `gorm:"not null" json:"nome"`
	Email string `gorm:"not null;unique" json:"email"`
}
