package iomodel

type Filter struct {
	Lat       *float64 `json:"lat"`
	Lng       *float64 `json:"lng"`
	VendorID  *int     `json:"vendor_id"`
	DateFrom  *string  `json:"date_from"`
	DateTo    *string  `json:"date_to"`
	Guest     *int     `json:"guest"`
	Bedroom   *int     `json:"bedroom"`
	NumBed    *int     `json:"num_bed"`
	PriceFrom *int     `json:"price_from"`
	PriceTo   *int     `json:"price_to"`
}
