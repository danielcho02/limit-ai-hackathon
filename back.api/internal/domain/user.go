package domain

type RoleType int

const (
	DefaultRole  RoleType = 0
	UserRole  RoleType = 1
	AdminRole RoleType = 999
)

type User struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Email    string   `json:"email" gorm:"unique;not null"`
	PasswordHash string   `json:"password" gorm:"not null"`
	Role     RoleType `json:"role" gorm:"type:enum('user', 'admin');default:'user'"`
}