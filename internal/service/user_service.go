package service
import (
	"context"
	"travel-web-backend/internal/entity"
	"travel-web-backend/internal/repository"
)
type UserService interface {
	Login(ctx context.Context, user model.User) (model.User, error)
	SignUp(ctx context.Context, user model.User) (model.User, error)
	ForgotPassword(ctx context.Context, user model.User) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(ctx context.Context, user model.User) (model.User, error) {
	return s.repo.Login(ctx, user)
}

func (s *userService) SignUp(ctx context.Context, user model.User) (model.User, error) {
	return s.repo.SignUp(ctx, user)
}

func (s *userService) ForgotPassword(ctx context.Context, user model.User) (model.User, error) {
	return s.repo.ForgotPassword(ctx, user)
}