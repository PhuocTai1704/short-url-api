package service

import "short-url-api/internal/repository"

type linkService struct {
	repo repository.LinkRepo
}

func NewLinkService(repo repository.LinkRepo) LinkService {
	return &linkService{
		repo: repo,
	}
}