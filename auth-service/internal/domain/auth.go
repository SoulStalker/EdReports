package domain

import (
	"time"

	"github.com/google/uuid"
)

// AuthKey - ключи подписи JWT
type AuthKey struct {
	Kid       string    `gorm:"primaryKey;type:text"` // key id
	Alg       string    `gorm:"not null;type:text"`   // 'EdDSA' или 'RS256'
	PublicPem string    `gorm:"not null;type:text"`   // публичный ключ PEM
	Status    string    `gorm:"not null;type:text"`   // 'active' | 'retired'
	CreatedAt time.Time `gorm:"default:now()"`
}

func (AuthKey) TableName() string {
	return "auth_keys"
}

// RefreshToken - refresh токены
type RefreshToken struct {
	ID            uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null"` // ссылка на пользователя
	SessionID     uuid.UUID  `gorm:"type:uuid;not null"` // постоянный id устройства/сессии
	TokenHash     string     `gorm:"not null;type:text"` // hash(refresh) = sha256/base64
	UserAgent     string     `gorm:"type:text"`          // user agent браузера
	IP            string     `gorm:"type:inet"`          // IP адрес
	CreatedAt     time.Time  `gorm:"default:now()"`      // время создания
	ExpiresAt     time.Time  `gorm:"not null"`           // время истечения
	RotatedFrom   *uuid.UUID `gorm:"type:uuid"`          // предыдущий токен этой сессии
	RevokedAt     *time.Time `gorm:"type:timestamptz"`   // время отзыва
	RevokedReason string     `gorm:"type:text"`          // причина отзыва

	// Связи
	// User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
