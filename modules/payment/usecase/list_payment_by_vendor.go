package paymentusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"

	"github.com/samber/lo"
)

func (uc *paymentUseCase) ListPaymentByVendorID(ctx context.Context, paging *common.Paging, vendorID int, bookingId int) ([]entities.Payment, error) {

	if paging.Limit == 0 {
		paging.Limit = constant.PaymentPagingLimitMax
		paging.Page = constant.PaymentPagingPageDefault
	}

	payments, err := uc.paymentSto.GetPaymentByVendor(ctx, int(vendorID), paging)
	if err != nil {
		return nil, err
	}

	if bookingId != 0 {
		res := lo.Filter(payments, func(item entities.Payment, _ int) bool {
			return item.BookingID == bookingId
		})
		return res, nil
	}

	return payments, nil
}
