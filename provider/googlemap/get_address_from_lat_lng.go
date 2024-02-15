package googlemapprovider

import (
	"context"
	"errors"
	"log"
	"paradise-booking/constant"
)

type ListAddressComponent []AddressComponent

func findAddress(lstAdressComponent []AddressComponent, res *GoogleMapAddress) {

	n := len(lstAdressComponent)
	limit := 3

	for i := n - 1; i >= 0; i-- {
		if limit == 0 {
			break
		}

		if limit == 3 {
			res.Country = lstAdressComponent[i].LongName
		} else if limit == 2 {
			res.State = lstAdressComponent[i].LongName
		} else if limit == 1 {
			res.District = lstAdressComponent[i].LongName
		}

		limit -= 1
	}
}

func (g *GoogleMap) GetAddressFromLatLng(ctx context.Context, lat, lng float64) (res *GoogleMapAddress, err error) {

	res = &GoogleMapAddress{}
	addressResp, err := g.GetGeocodeMap(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	if len(addressResp.Results) == 0 {
		log.Println("err", "not found any address")
		return nil, errors.New("not found any address")
	}

	//ROOFTOP first; RANGE_INTERPOLATED second, GEOMETRIC_CENTER third, APPROXIMATE last
	mapLocationTypeAnd := map[string]ListAddressComponent{}
	for _, address := range addressResp.Results {
		mapLocationTypeAnd[address.Geometry.LocationType] = address.AddressComponents
	}
	// case1 : type = ROOFTOP
	if addressComponents, ok := mapLocationTypeAnd[constant.LocationType_Rooftop]; ok {
		findAddress(addressComponents, res)
		return res, nil
	} else if addressComponents, ok := mapLocationTypeAnd[constant.LocationType_RangeInterpolated]; ok {
		findAddress(addressComponents, res)
		return res, nil
	} else if addressComponents, ok := mapLocationTypeAnd[constant.LocationType_GeometricCenter]; ok {
		findAddress(addressComponents, res)
		return res, nil
	} else {
		findAddress(addressComponents, res)
		return res, nil
	}
}
