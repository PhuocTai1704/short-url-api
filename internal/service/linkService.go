package service

import (
	"os"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"short-url-api/internal/repository"
	"short-url-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type linkService struct {
	repo repository.LinkRepo
}

func NewLinkService(repo repository.LinkRepo) LinkService {
	return &linkService{
		repo: repo,
	}
}

func (lk *linkService) CreateLink(ctx *gin.Context, url string) (dto.LinkDTO, error){	
	
	domain := os.Getenv("DOMAIN_SHORT")

    link := model.Link{
        Url:  url,
		Code: utils.DeterministicShort(url,8),
    }

	if err := lk.repo.FirstOrCreate(&link); err != nil {
		return dto.LinkDTO{}, err
	}
	return dto.LinkDTO{
		ID:        link.ID,
		Url:       link.Url,
		Code:      link.Code,
		Link: 	   domain+link.Code,
		Clicks:    link.Clicks,
		LastClick: link.LastClick,
		CreatedAt: link.CreatedAt,
	}, nil
}