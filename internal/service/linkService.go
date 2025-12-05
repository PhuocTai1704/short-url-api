package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"short-url-api/internal/repository"
	"short-url-api/internal/utils"

	"gorm.io/gorm"
)

type linkService struct {
	repo repository.LinkRepo
}

func NewLinkService(repo repository.LinkRepo) LinkService {
	return &linkService{
		repo: repo,
	}
}

// Tạo link với alias do user truyền
func (lk *linkService) createWithAlias(ctx context.Context, link *model.Link, alias, domain string) error {
    if len(alias) < 5 {
        return fmt.Errorf("alias phải có ít nhất 5 ký tự")
    }

    exists, err := lk.repo.IsCodeExist(ctx, alias)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("alias '%s' đã tồn tại", alias)
    }

    link.Code = alias
    link.Link = domain + alias

    return lk.repo.Create(ctx, link)
}
// Tạo link với code tự sinh
func (lk *linkService) createAutoCode(ctx context.Context, link *model.Link, url, domain string) error {
    k := 8
    code := utils.DeterministicShort(url, k)
    maxAttempts := 10

    for i := 0; i < maxAttempts; i++ {
        linkCheck, err := lk.repo.GetByCode(ctx, code)
        if err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                link.Code = code
                link.Link = domain + code
                return lk.repo.Create(ctx, link)
            }
            return err
        }

        if linkCheck.Url == url {
            *link = *linkCheck
            return nil
        }

        k++
        code = utils.DeterministicShort(url, k)
    }

    return fmt.Errorf("không tạo được short code sau %d lần thử", maxAttempts)
}

func (lk *linkService) CreateLink(ctx context.Context, url, alias string) (dto.LinkDTO, error) {
    domain := os.Getenv("DOMAIN_SHORT")
    var link model.Link
    link.Url = url

    if alias != "" {
        if err := lk.createWithAlias(ctx, &link, alias, domain); err != nil {
            return dto.LinkDTO{}, err
        }
    } else {
        if err := lk.createAutoCode(ctx, &link, url, domain); err != nil {
            return dto.LinkDTO{}, err
        }
    }

    // Chỉ return 1 lần ở cuối
    return dto.LinkDTO{
        ID:        link.ID,
        Url:       link.Url,
        Code:      link.Code,
        Link:      link.Link,
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
