package placeusecase

import (
	"context"
	"iTask/common"
	"iTask/config"
	"iTask/entities"
	"iTask/modules/place/iomodel"
	googlemapprovider "iTask/provider/googlemap"
	"time"
)

type PlaceStorage interface {
	Create(ctx context.Context, data *entities.Place) (err error)
	DeleteByID(ctx context.Context, id int) error
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceByVendorID(ctx context.Context, vendorID int, paging *common.Paging) ([]entities.Place, error)
	ListPlaces(ctx context.Context, paging *common.Paging, filter *iomodel.Filter, address *googlemapprovider.GoogleMapAddress) ([]entities.Place, error)
	UpdateByID(ctx context.Context, id int, data *entities.Place) error
	GetPlaceByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.Place, error)
	GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error)
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type PlaceStoCache interface {
	GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error)
}

type PlaceWishListSto interface {
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error)
}

type BookingSto interface {
	ListAllBookingWithCondition(ctx context.Context, condition []common.Condition) ([]entities.Booking, error)
	GetBookingsWithinDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]entities.Booking, error)
}

type placeUseCase struct {
	placeStorage  PlaceStorage
	accountSto    AccountStorage
	placeWishSto  PlaceWishListSto
	cfg           *config.Config
	googleMap     *googlemapprovider.GoogleMap
	placeStoCache PlaceStoCache
	bookingSto    BookingSto
}

func NewPlaceUseCase(cfg *config.Config, placeSto PlaceStorage, accoutSto AccountStorage, googleMap *googlemapprovider.GoogleMap, placeWishSto PlaceWishListSto, placeStoCache PlaceStoCache, bookingSto BookingSto) *placeUseCase {
	return &placeUseCase{placeSto, accoutSto, placeWishSto, cfg, googleMap, placeStoCache, bookingSto}
}
