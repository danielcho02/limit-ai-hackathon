package domain

type RoleType int

const (
	DefaultRole  RoleType = 0
	UserRole  RoleType = 1
	OrganizationRole RoleType = 2
	AdminRole RoleType = 999
)

type User struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Nickname string 	`json:"nickname"`
	PasswordHash string   `json:"password" gorm:"not null"`
	Role     RoleType `json:"role" gorm:"type:enum('user', 'admin');default:'user'"`
}