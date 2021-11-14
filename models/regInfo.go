package models

type RegInfo struct {
	ID       int64  `gorm:"id" json:"id" `
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"password" json:"password"`
	Code     string  `gorm:"code" json:"code"`
}

type RegUserInfoArgs struct {

	Email    string `gorm:"email;primary_key;unique" json:"email" binding:"omitempty,email" validate:"required"`
	Code     string  `gorm:"code" json:"code" binding:"omitempty,code" validate:"required" `
	Password string  `gorm:"password" json:"password" binding:"omitempty,password" validate:"required"`
}
type RegUserInfoExceptCodeArgs struct {
	Email    string `gorm:"email;primary_key;unique" json:"email" binding:"omitempty,email" validate:"required"`

	Password string  `gorm:"password" json:"password" binding:"omitempty,password" validate:"required"`
}

