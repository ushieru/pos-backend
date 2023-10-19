package dto

import "github.com/ushieru/pos/domain"

type AuthUserResponse struct {
	Token string      `json:"token"`
	User  domain.User `json:"user"`
}
