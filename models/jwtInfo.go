package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims   // playload
}

type Reg struct {
	ID int `gorm:"primary_key" json:"id"`
	Email string `json:"username"`
	Password string `json:"password"`
	Code int `json:"code"`
}