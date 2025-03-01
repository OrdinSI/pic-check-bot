package model

import "time"

type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	HashPart1 uint64    `json:"hash_part1" gorm:"column:hash_part1;not null"`
	HashPart2 uint64    `json:"hash_part2" gorm:"column:hash_part2;not null"`
	FileHash  []byte    `json:"file_hash" gorm:"not null;type:BYTEA"`
	UserID    int64     `json:"user_id" gorm:"not null"`
	GroupID   int64     `json:"group_id" gorm:"not null"`
	FileID    string    `json:"file_id" gorm:"not null"`
	MessageID int       `json:"message_id" gorm:"not null"`
	PostTime  time.Time `json:"post_time" gorm:"default:CURRENT_TIMESTAMP"`
}

func (i *Image) TableName() string {
	return "images"
}

type Repost struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageID    uint      `json:"image_id" gorm:"not null"`
	UserID     int64     `json:"user_id" gorm:"not null"`
	GroupID    int64     `json:"group_id" gorm:"not null"`
	RepostTime time.Time `json:"repost_time" gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Repost) TableName() string {
	return "reposts"
}

type TopRepost struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Count    int    `json:"count"`
}
