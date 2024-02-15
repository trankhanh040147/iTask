package bookingusecase

import (
	"context"
	"iTask/common"
	"iTask/config"
	"iTask/entities"
	"iTask/modules/booking/iomodel"
	bookingdetailstorage "iTask/modules/booking_detail/storage"
	momoprovider "iTask/provider/momo"
	"iTask/worker"
	"time"
)

type BookingStorage interface {
	Create(ctx context.Context, data *entities.Booking) (err error)
	UpdateStatus(ctx context.Context, bookingID int, status int) error
	ListByFilter(ctx context.Context, filter *iomodel.FilterListBooking, paging *common.Paging, userId int) ([]entities.Booking, error)
	GetByID(ctx context.Context, id int) (*entities.Booking, error)
	GetByPlaceID(ctx context.Context, placeId int, paging *common.Paging) ([]entities.Booking, error)
	ListPlaceIds(ctx context.Context) ([]int, error)
	ListAllBookingWithCondition(ctx context.Context, condition []common.Condition) ([]entities.Booking, error)
	DeleteByID(ctx context.Context, id int) error
	GetBookingsWithinDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]entities.Booking, error)
}

type BookingDetailStorage interface {
	Create(ctx context.Context, data *entities.BookingDetail) (err error)
	CreateTx(ctx context.Context, createBookingDetailTxParam bookingdetailstorage.CreateBookingDetailTxParam) error
	GetByBookingID(ctx context.Context, bookingId int) (res *entities.BookingDetail, err error)
}

type AccountSto interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type PlaceSto interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceNotInIds(ctx context.Context, placeIds []int, vendorId int) ([]entities.Place, error)
	ListPlaceInIds(ctx context.Context, placeIds []int, vendorId int) ([]entities.Place, error)
	UpdateWithMap(ctx context.Context, place *entities.Place, props map[string]interface{}) error
	ListPlaceByCondition(ctx context.Context, condition []common.Condition) ([]entities.Place, error)
}

type PaymentSto interface {
	CreatePayment(ctx context.Context, payment *entities.Payment) error
}

type bookingUseCase struct {
	bookingSto       BookingStorage
	bookingDetailSto BookingDetailStorage
	AccountSto       AccountSto
	paymentSto       PaymentSto
	cfg              *config.Config
	taskDistributor  worker.TaskDistributor
	PlaceSto         PlaceSto
	MomoProvider     *momoprovider.Momo
}

func NewBookingUseCase(bookingStore BookingStorage, bookingDetailStorage BookingDetailStorage, config *config.Config, taskDistributor worker.TaskDistributor, accountSto AccountSto, placeSto PlaceSto, momo *momoprovider.Momo, paymentSto PaymentSto) *bookingUseCase {
	return &bookingUseCase{bookingSto: bookingStore, bookingDetailSto: bookingDetailStorage, cfg: config, taskDistributor: taskDistributor, AccountSto: accountSto, PlaceSto: placeSto, MomoProvider: momo, paymentSto: paymentSto}
}
