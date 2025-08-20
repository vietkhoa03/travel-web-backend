package service

import (
	"context"
	model "travel-web-backend/internal/entity"
	repository "travel-web-backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type LocationService interface {
	GetAllLocation(ctx context.Context) ([]model.Location, error)
	CreateLocation(ctx context.Context, location model.Location) (model.Location, error)
	UpdateLocation(ctx context.Context, id string, location model.Location) (model.Location, error)
	SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error)
	GetLocationByID(ctx context.Context, id string) (model.Location, error)
	DeleteLocation(ctx context.Context, id string) error
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

func (s *locationService) GetAllLocation(ctx context.Context) ([]model.Location, error) {
	return s.repo.GetAllLocation(ctx)
}

func (s *locationService) GetLocationByID(ctx context.Context, id string) (model.Location, error) {
	return s.repo.GetLocationByID(ctx, id)
}

func (s *locationService) CreateLocation(ctx context.Context, location model.Location) (model.Location, error) {
	return s.repo.CreateLocation(ctx, location)
}

func (s *locationService) UpdateLocation(ctx context.Context, id string, location model.Location) (model.Location, error) {
	return s.repo.UpdateLocation(ctx, id, location)
}

func (s *locationService) DeleteLocation(ctx context.Context, id string) error {
	return s.repo.DeleteLocation(ctx, id)
}
