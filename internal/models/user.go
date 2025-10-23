package models

import (
	"time"
)

type Role string

const (
	RoleUser           Role = "user"
	RoleContentCreator Role = "content_creator"
	RoleAdmin          Role = "admin"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TrackmaniaAccountID   string `json:"trackmania_account_id" gorm:"uniqueIndex;not null"`
	TrackmaniaUsername    string `json:"trackmania_username"`
	TrackmaniaDisplayName string `json:"trackmania_display_name"`

	AccessToken  string    `json:"-" gorm:"type:text"`
	RefreshToken string    `json:"-" gorm:"type:text"`
	TokenExpiry  time.Time `json:"-"`
	AvatarURL    string    `json:"avatar_url"`
	Role         Role      `json:"role" gorm:"type:varchar(50);default:'user';not null"`

	LastLoginAt *time.Time `json:"last_login_at"`
}

func (u *User) IsTokenExpired() bool {
	return time.Now().After(u.TokenExpiry)
}

func (u *User) NeedsTokenRefresh() bool {
	return time.Now().Add(5 * time.Minute).After(u.TokenExpiry)
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsContentCreator() bool {
	return u.Role == RoleContentCreator || u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser || u.Role == RoleContentCreator || u.Role == RoleAdmin
}

func (u *User) HasRole(role Role) bool {
	roleHierarchy := map[Role]int{
		RoleUser:           1,
		RoleContentCreator: 2,
		RoleAdmin:          3,
	}

	return roleHierarchy[u.Role] >= roleHierarchy[role]
}
