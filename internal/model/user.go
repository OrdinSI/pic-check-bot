package model

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at" gorm:"->;<-:create;"`
	Images    []Image   `json:"images" gorm:"foreignKey:UserID"`
	Reposts   []Repost  `json:"reposts" gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}

type Group struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	GroupName string    `json:"group_name"`
	CreatedAt time.Time `json:"created_at" gorm:"->;<-:create;"`
}

func (g *Group) TableName() string {
	return "groups"
}
