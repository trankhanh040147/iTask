package googlemapprovider

type GoogleMapAddress struct {
	Country  string `json:"country"`
	State    string `json:"state"`
	District string `json:"district"`
}

type GoogleMapResponse struct {
	PlusCodeData PlusCode `json:"plus_code"`
	Results      []Result `json:"results"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type Result struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	PlaceID           string             `json:"place_id"`
	Geometry          Geometry           `json:"geometry"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	LocationType string `json:"location_type"`
}
