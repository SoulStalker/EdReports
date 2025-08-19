package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User - пользователь
type User struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email        string    `gorm:"uniqueIndex;not null;type:text"`         // email
	PasswordHash string    `gorm:"not null;type:text"`                     // bcrypt/argon2id хеш
	Name         string    `gorm:"type:text"`                              // имя
	Active       bool      `gorm:"default:true"`                           // активен ли
	Tz           string    `gorm:"default:'Asia/Yekaterinburg';type:text"` // часовой пояс
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"default:now()"`

	// Связи
	Roles         []Role         `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
	Stores        []UserStore    `gorm:"constraint:OnDelete:CASCADE"`
	Categories    []UserCategory `gorm:"constraint:OnDelete:CASCADE"`
	RefreshTokens []RefreshToken `gorm:"constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}

// Хук для обновления updated_at
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
