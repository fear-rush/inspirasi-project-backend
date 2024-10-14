package model

type EarthquakeData struct {
	ID        string `json:"_id"`
	Latitude  string `json:"lintang"`
	Longitude string `json:"bujur"`
	Magnitude string `json:"magnitude"`
	Depth     string `json:"kedalaman"`
	Region    string `json:"wilayah"`
	DateTime  string `json:"datetime"`
}
