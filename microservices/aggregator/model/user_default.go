package model

type UserDefault struct {
	UserUuid string `json:"user_uuid"`
	UserName string `json:"user_name,omitempty"`
}
