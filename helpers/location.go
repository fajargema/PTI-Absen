package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLocation(latitude, longitude string) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%s&lon=%s", latitude, longitude)
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
