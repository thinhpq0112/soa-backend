package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	lat1, lon1 := 48.8588897, 2.32      // Paris
	lat2, lon2 := 45.7578137, 4.8320114 // Lyon

	expectedDistance := 465.07 // Approximate distance in km

	distance := CalculateDistance(lat1, lon1, lat2, lon2)
	assert.InDelta(t, expectedDistance, distance, 80.0, "Expected distance to be around %.2f km, but got %.2f km", expectedDistance, distance)
}

func TestGetLatLonFromIP(t *testing.T) {
	tests := []struct {
		ip          string
		expectedLat float64
		expectedLon float64
	}{
		{"8.8.8.8", 39.030, -77.500},
		{"222.253.48.127", 10.7826337, 106.7693117},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			lat, lon, err := GetLatLonFromIP(tt.ip)
			assert.NoError(t, err, "Expected no error")
			assert.InDelta(t, tt.expectedLat, lat, 1.0, "Expected latitude to be around %.3f, but got %.3f", tt.expectedLat, lat)
			assert.InDelta(t, tt.expectedLon, lon, 1.0, "Expected longitude to be around %.3f, but got %.3f", tt.expectedLon, lon)
		})
	}
}

func TestGetLatLonFromCity(t *testing.T) {
	tests := []struct {
		city        string
		expectedLat float64
		expectedLon float64
	}{
		{"Paris", 48.8588897, 2.32},
		{"Lyon", 45.7578137, 4.8320114},
	}

	for _, tt := range tests {
		t.Run(tt.city, func(t *testing.T) {
			lat, lon, err := GetLatLonFromCity(tt.city)
			assert.NoError(t, err, "Expected no error")
			assert.InDelta(t, tt.expectedLat, lat, 1.0, "Expected latitude to be around %.4f, but got %.4f", tt.expectedLat, lat)
			assert.InDelta(t, tt.expectedLon, lon, 1.0, "Expected longitude to be around %.4f, but got %.4f", tt.expectedLon, lon)
		})
	}
}
