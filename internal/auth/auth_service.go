package auth

import "context"

type Service struct {
	repo *Repository
	jwt  *JWTManager
}

func NewService(repo *Repository, jwt *JWTManager) *Service {
	return &Service{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	hash, err := HashPassword(password)
	if err != nil {
		return "", err
	}

	id, err := s.repo.CreateUser(ctx, email, hash)
	if err != nil {
		return "", err
	}

	return s.jwt.Generate(id)
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	id, hash, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := CheckPassword(password, hash); err != nil {
		return "", err
	}

	return s.jwt.Generate(id)
}
