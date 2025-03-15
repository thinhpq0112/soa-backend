package service

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/util"
)

type DistanceService struct {
	//repo repository.IProductRepo
}

func NewDistanceService() *DistanceService {
	return &DistanceService{}
}

func (s *DistanceService) CalculateDistance(ctx context.Context, ip, cityName string) (float64, error) {
	userLat, userLon, err := util.GetLatLonFromIP(ip)
	if err != nil {
		return 0, err
	}

	cityLat, cityLon, err := util.GetLatLonFromCity(cityName)
	if err != nil {
		return 0, err
	}

	return util.CalculateDistance(userLat, userLon, cityLat, cityLon), nil
}
