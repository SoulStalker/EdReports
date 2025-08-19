package domain

import "github.com/google/uuid"

// Role - роль
type Role struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex;not null;type:text"` // 'admin', 'ceo', etc.
	Description string    `gorm:"type:text"`                      // описание

	// Связи
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE"`
	Users       []User       `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}

func (Role) TableName() string {
	return "roles"
}

// Permission - право доступа
type Permission struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Code        string    `gorm:"uniqueIndex;not null;type:text"` // 'reports:read', 'users:write'
	Description string    `gorm:"type:text"`                      // описание

	// Связи
	Roles []Role `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE"`
}

func (Permission) TableName() string {
	return "permissions"
}

// RolePermission - связь ролей и прав (many-to-many)
type RolePermission struct {
	RoleID       uuid.UUID `gorm:"primaryKey;type:uuid"`
	PermissionID uuid.UUID `gorm:"primaryKey;type:uuid"`

	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

// UserRole - связь пользователей и ролей (many-to-many)
type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey;type:uuid"`
	RoleID uuid.UUID `gorm:"primaryKey;type:uuid"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
