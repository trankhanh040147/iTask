package googlemapprovider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"paradise-booking/config"
	"paradise-booking/utils"
	"testing"

	"github.com/spf13/viper"
)

func LoadConfigTest() (*config.Config, error) {
	v := viper.New()

	v.AddConfigPath("../../config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// check is not found file config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c config.Config // Unmarshal data config have get in file config then get into c
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func TestGetGeocodeMap(t *testing.T) {
	cfg, _ := LoadConfigTest()
	g := NewGoogleMap(cfg)
	lat := 13.883407
	lng := 109.119524
	latLngValue := fmt.Sprintf("%f,%f", lat, lng)
	var resp *GoogleMapResponse
	path := fmt.Sprintf("/json?latlng=%s&key=%s", latLngValue, g.cfg.GoogleMap.APIKey)
	url := utils.JoinURL(g.cfg.GoogleMap.BaseURL, path)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Println("err", err)
	}

	body, err := g.MakeRequest(context.Background(), g.client, req)
	if err != nil {
		log.Println("err", err)
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		log.Println("err", err)
	}

	log.Println(body)
}

// trigger cicd
