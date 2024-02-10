package entities

import "time"

type VerifyEmail struct {
	Id        int        `json:"id" gorm:"column:id"`
	Email     string     `json:"email" gorm:"column:email"`
	ScretCode string     `json:"scret_code" gorm:"column:scret_code"`
	Type      int        `json:"type" gorm:"column:type"` // 1: verify email, 2: reset password
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	ExpiredAt *time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func (VerifyEmail) TableName() string {
	return "verify_emails"
}

func (v *VerifyEmail) IsExpired() bool {
	return v.ExpiredAt.Before(time.Now())
}
