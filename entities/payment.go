package entities

import "paradise-booking/common"

type Payment struct {
	common.SQLModel
	BookingID int     `json:"booking_id" gorm:"column:booking_id"`
	MethodID  int     `json:"method_id" gorm:"column:method_id"`
	StatusID  int     `json:"status_id" gorm:"column:status_id"`
	Amount    float64 `json:"amount" gorm:"column:amount"`
	RequestID string  `json:"request_id" gorm:"column:request_id"`
	OrderID   string  `json:"order_id" gorm:"column:order_id"`
}

func (Payment) TableName() string {
	return "payments"
}
