package amenityhandler

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/amenity/iomodel"
)

type AmenityUseCase interface {
	CreateAmenity(ctx context.Context, data *iomodel.CreateAmenityReq) (err error)
	DeleteAmenityById(ctx context.Context, id int) error
	GetAllConfigAmenity(ctx context.Context) (res []entities.ConfigAmenity, err error)
	ListAmenityByPlaceID(ctx context.Context, placeID int) (res []entities.Amenity, err error)
	DeleteAmenityByListId(ctx context.Context, req *iomodel.DeleteAmenityReq) error
}

type amenityHandler struct {
	amenityUC AmenityUseCase
}

func NewAmenityHandler(amenity AmenityUseCase) *amenityHandler {
	return &amenityHandler{amenityUC: amenity}
}
