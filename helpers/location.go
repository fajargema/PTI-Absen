package helpers

import (
	"absen/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	GOOGLE_API_KEY string = utils.GetConfig("GOOGLE_API_KEY")
)

func GetLocation(latitude, longitude string) (string, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", latitude, longitude, GOOGLE_API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
