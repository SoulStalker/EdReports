package models

// User - пользователь
type User struct {
	ID           int64  `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex;not null;type:text"` // email
	PasswordHash []byte `gorm:"not null;type:text"`             // bcrypt/argon2id хеш
	// Name         string    `gorm:"type:text"`                            // имя
	// Active       bool      `gorm:"default:true"`                         // активен ли
	// Tz           string    `gorm:"default:Asia/Yekaterinburg;type:text"` // часовой пояс
	// CreatedAt    time.Time `gorm:"default:now()"`
	// UpdatedAt    time.Time `gorm:"default:now()"`

	// Связи
	// Roles         []Role         `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
	// Stores        []UserStore    `gorm:"constraint:OnDelete:CASCADE"`
	// Categories    []UserCategory `gorm:"constraint:OnDelete:CASCADE"`
	// RefreshTokens []RefreshToken `gorm:"constraint:OnDelete:CASCADE"`
}

// func (User) TableName() string {
// 	return "users"
// }

// Хук для обновления updated_at
// func (u *User) BeforeUpdate(tx *gorm.DB) error {
// 	u.UpdatedAt = time.Now()
// 	return nil
// }

// type UserRepository interface {
// 	Create(user *User) error
// 	GetByEmail(email string) (*User, error)
// 	GetByID(id string) (*User, error)
// 	AssignRole(userID, roleID string) error
// }

// type PasswordHasher interface {
// 	Hash(password string) (string, error)
// 	Compare(hash, password string) bool
// }

// type TokenProvider interface {
// 	GenerateAccessToken(user *User) (string, error)
// 	GenerateRefreshToken(user *User) (string, error)
// 	ValidateToken(token string) (*User, error)
// }
