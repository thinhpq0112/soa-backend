package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jftuga/geodist"
	"net/http"
	"strconv"
	"strings"
)

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	point1 := geodist.Coord{Lat: lat1, Lon: lon1}
	point2 := geodist.Coord{Lat: lat2, Lon: lon2}

	_, kiloM := geodist.HaversineDistance(point1, point2)
	return kiloM
}

type IPGeoResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GeoResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetLatLonFromIP(ip string) (float64, float64, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var result IPGeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}

	return result.Lat, result.Lon, nil
}

func GetLatLonFromCity(city string) (float64, float64, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", strings.ReplaceAll(city, " ", "+"))
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var results []GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, err
	}

	if len(results) == 0 {
		return 0, 0, errors.New("city not found")
	}

	lat, lon := results[0].Lat, results[0].Lon
	latF, _ := strconv.ParseFloat(lat, 64)
	lonF, _ := strconv.ParseFloat(lon, 64)
	return latF, lonF, nil
}
