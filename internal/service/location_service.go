package service

import (
	"context"
	"travel-web-backend/internal/entity"
	repository "travel-web-backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type LocationService interface {
	SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error)
	GetLocationByID(ctx context.Context, id string) (model.Location, error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error) {
	filter := bson.M{}
	if query != "" {
		filter["name"] = bson.M{
			"$regex":   query,
			"$options": "i",
		}
	}
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return s.repo.SearchLocationsByName(ctx, query, page, limit)
}

func (s *locationService) GetLocationByID(ctx context.Context, id string) (model.Location, error) {
	return s.repo.GetLocationByID(ctx, id)
}
