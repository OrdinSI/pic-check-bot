package model

type RequestStart struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

type UserRequest struct {
	UserID    int64  `json:"user_id"`
	GroupID   int64  `json:"group_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ImageRequest struct {
	FileID    string `json:"file_id"`
	UserID    int64  `json:"user_id"`
	GroupID   int64  `json:"group_id"`
	MessageID int    `json:"message_id"`
}
