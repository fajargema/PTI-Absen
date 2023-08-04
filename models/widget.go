package models

type HomeWidget struct {
	CheckIn  CheckIn  `json:"check_in"`
	CheckOut CheckOut `json:"check_out"`
}

type CheckIn struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Time      string `json:"time"`
	Location  string `json:"location"`
}

type CheckOut struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Time      string `json:"time"`
	Location  string `json:"location"`
}

type LocationDetails struct {
	DisplayName string `json:"display_name"`
}
