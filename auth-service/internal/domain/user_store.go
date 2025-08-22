package domain

// import "github.com/google/uuid"

// // UserStore - привязка пользователя к магазину
// type UserStore struct {
// 	UserID  uuid.UUID `gorm:"primaryKey;type:uuid"`
// 	StoreID uuid.UUID `gorm:"primaryKey;type:uuid"` // id из CatalogService

// 	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
// }

// func (UserStore) TableName() string {
// 	return "user_stores"
// }

// // UserCategory - привязка пользователя к категории
// type UserCategory struct {
// 	UserID     uuid.UUID `gorm:"primaryKey;type:uuid"`
// 	CategoryID uuid.UUID `gorm:"primaryKey;type:uuid"` // id из CatalogService

// 	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
// }

// func (UserCategory) TableName() string {
// 	return "user_categories"
// }
