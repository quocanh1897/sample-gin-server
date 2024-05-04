package model

import "github.com/google/uuid"

type Authorization struct {
	Jwt                string
	Scopes             []string
	OauthApplicationId *uuid.UUID
}
