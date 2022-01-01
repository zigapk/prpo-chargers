package gmaps

import (
	"encoding/json"
	"fmt"
	"github.com/zigapk/prpo-chargers/internal/config"
	"io/ioutil"
	"net/http"
)

type gmapsResponse struct {
	Results []gmapsResult `json:"results"`
}

type gmapsResult struct {
	FormattedAddress string `json:"formatted_address"`
}

func GetAddressFromCoordinates(lat float64, lon float64) (string, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=%s", lat, lon, config.Server.GmapsApiKey())
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	gResponse := &gmapsResponse{}
	err = json.Unmarshal(body, gResponse)
	if err != nil {
		return "", err
	}

	return gResponse.Results[0].FormattedAddress, nil
}
