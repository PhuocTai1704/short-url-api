package service

import (
	"context"
	"os"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"short-url-api/internal/repository"
	"short-url-api/internal/utils"
)

type linkService struct {
	repo repository.LinkRepo
}

func NewLinkService(repo repository.LinkRepo) LinkService {
	return &linkService{
		repo: repo,
	}
}

func (lk *linkService) CreateLink(ctx context.Context, url string) (dto.LinkDTO, error){	
	domain := os.Getenv("DOMAIN_SHORT")
	code := utils.DeterministicShort(url,8)
    link := model.Link{
        Url:  url,
		Code: code,
		Link: domain+code,
    }

	if err := lk.repo.FirstOrCreate(ctx,&link); err != nil {
		return dto.LinkDTO{}, err
	}
	return dto.LinkDTO{
		ID:        link.ID,
		Url:       link.Url,
		Code:      link.Code,
		Link: 	   link.Link,
		Clicks:    link.Clicks,
		LastClick: link.LastClick,
		CreatedAt: link.CreatedAt,
	}, nil
}

func (s *linkService) GetAllLinks(ctx context.Context, page, limit int) ([]dto.LinkDTO, int64, bool, error) {

    links, total, err := s.repo.GetAllLinks(ctx, page, limit)
    if err != nil {
        return nil, 0,true, err
    }

    // Convert sang DTO
    linkDTOs := make([]dto.LinkDTO, len(links))
    for i, l := range links {
        linkDTOs[i] = dto.LinkDTO{
            ID:      l.ID,
            Url:     l.Url,
            Code:    l.Code,
			Link:    l.Link,
            CreatedAt: l.CreatedAt,
        }
    }
    isLast := page*limit >= int(total)

    return linkDTOs, total, isLast, nil
}
