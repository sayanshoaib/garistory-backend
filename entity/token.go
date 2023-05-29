package entity

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	Id                string `json:"id"`
	AccountType       string `json:"account_type"`
	AccountCategoryId int    `json:"account_category_id"`
}
type JwtClaims struct {
	UserId      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	AccountType int    `json:"account_type"`
	jwt.StandardClaims
}
