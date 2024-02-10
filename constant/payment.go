package constant

const (
	PaymentMethodCod = iota + 1
	PaymentMethodMomo
)

const (
	PaymentStatusUnpaid = iota + 1
	PaymentStatusPaid
)

const (
	PaymentPagingLimitMax    = 200
	PaymentPagingPageDefault = 1
)
