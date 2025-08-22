package domain

// import (
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // CreateUserRequest - DTO для создания пользователя
// // todo наверно dto надо будет в другое место перенести и разнести все из файла в другие файлы
// type CreateUserRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8"`
// 	Name     string `json:"name" validate:"required"`
// 	Tz       string `json:"tz" validate:"required"`
// }

// // LoginRequest - DTO для входа
// type LoginRequest struct {
// 	Email     string `json:"email" validate:"required,email"`
// 	Password  string `json:"password" validate:"required"`
// 	UserAgent string `json:"user_agent"`
// 	IP        string `json:"ip"`
// }

// // TokenResponse - DTO для ответа с токенами
// type TokenResponse struct {
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// 	ExpiresIn    int64  `json:"expires_in"`
// }

// // GenerateSessionID генерирует новый session_id
// func GenerateSessionID() uuid.UUID {
// 	return uuid.New()
// }

// // BeforeCreate - хук для генерации ID
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.ID == uuid.Nil {
// 		u.ID = uuid.New()
// 	}
// 	return nil
// }

// func (r *Role) BeforeCreate(tx *gorm.DB) error {
// 	if r.ID == uuid.Nil {
// 		r.ID = uuid.New()
// 	}
// 	return nil
// }

// func (p *Permission) BeforeCreate(tx *gorm.DB) error {
// 	if p.ID == uuid.Nil {
// 		p.ID = uuid.New()
// 	}
// 	return nil
// }
