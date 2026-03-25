package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"short-url-api/internal/repository"
	"short-url-api/internal/utils"
	"time"

	"github.com/google/uuid"
)

type linkService struct {
	repo         repository.LinkRepo
	linkClickSvc LinkClickService
}

func NewLinkService(repo repository.LinkRepo, linkClickSvc LinkClickService) LinkService {
	return &linkService{
		repo:         repo,
		linkClickSvc: linkClickSvc,
	}
}

// Tạo link với alias do user truyền
func (lk *linkService) createWithAlias(ctx context.Context, link *model.Link, alias, baseUrl string) error {
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
	link.Link = baseUrl + alias

	return lk.repo.Create(ctx, link)
}

// Tạo link với code tự sinh
func (lk *linkService) createAutoCode(ctx context.Context, link *model.Link, baseUrl string) error {
	k := 8

	for attempts := 0; attempts < 5; attempts++ {

		code, err := utils.GenerateShortCode(k)
		if err != nil {
			return err
		}

		exists, err := lk.repo.ExistsByCode(ctx, code)
		if err != nil {
			return err
		}

		// Nếu code chưa tồn tại -> Tạo link
		if !exists {
			link.Code = code
			link.Link = baseUrl + code
			return lk.repo.Create(ctx, link)
		}

	}

	return fmt.Errorf("tạo short code thất bại: không tìm được mã hợp lệ")
}

func (lk *linkService) CreateLink(ctx context.Context, url, alias string) (dto.LinkDTO, error) {
	baseUrl := fmt.Sprintf("%s://%s:/", os.Getenv("PROTOCOL"), os.Getenv("DOMAIN_SHORT"))
	if !utils.ValidateURL(url) {
		return dto.LinkDTO{}, fmt.Errorf("url không hợp lệ")
	}
	var link model.Link
	link.Url = url

	if alias != "" {
		if err := lk.createWithAlias(ctx, &link, alias, baseUrl); err != nil {
			return dto.LinkDTO{}, err
		}
	} else {
		if err := lk.createAutoCode(ctx, &link, baseUrl); err != nil {
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
func (lk *linkService) UpdateLink(ctx context.Context, linkId uuid.UUID, url string, alias string) (dto.LinkDTO, error) {
	baseUrl := fmt.Sprintf("%s://%s/", os.Getenv("PROTOCOL"), os.Getenv("DOMAIN_SHORT"))
	if !utils.ValidateURL(url) {
		return dto.LinkDTO{}, fmt.Errorf("url không hợp lệ")
	}
	if len(alias) < 5 {
		return dto.LinkDTO{}, fmt.Errorf("alias phải có ít nhất 5 ký tự")
	}
	exists, err := lk.repo.IsCodeExistNotId(ctx, linkId, alias)
	if err != nil {
		return dto.LinkDTO{}, err
	}
	if exists {
		return dto.LinkDTO{}, fmt.Errorf("alias '%s' đã tồn tại", alias)
	}
	link := model.Link{
		ID:   linkId,
		Url:  url,
		Code: alias,
		Link: baseUrl + alias,
	}
	err = lk.repo.Update(ctx, &link)
	if err != nil {
		return dto.LinkDTO{}, err
	}

	// Chỉ return 1 lần ở cuối
	return dto.LinkDTO{
		ID:        link.ID,
		Url:       link.Url,
		Code:      link.Code,
		Link:      link.Link,
		CreatedAt: link.CreatedAt,
	}, nil
}

func (lk *linkService) GetById(ctx context.Context, id uuid.UUID) (dto.LinkDTO, error) {
	link, err := lk.repo.GetLinkWithStatsById(ctx, id)
	if err != nil {
		return dto.LinkDTO{}, err
	}

	return *link, nil
}

func (lk *linkService) GetByLink(ctx context.Context, urlLink string) (dto.LinkDTO, error) {
	link, err := lk.repo.GetLinkWithStatsByLink(ctx, urlLink)
	if err != nil {
		return dto.LinkDTO{}, err
	}

	return *link, nil
}

func (lk *linkService) GetUrlByCode(ctx context.Context, code string, req *http.Request) (string, error) {
	link, err := lk.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}

	// Tăng số lượt click + cập nhật last_click
	err = lk.linkClickSvc.IncreaseClick(ctx, link.ID, req)
	if err != nil {
		return "", err
	}

	return link.Url, nil
}

func (lk *linkService) GetAllLinks(ctx context.Context, page, limit int, startDate, endDate *time.Time) ([]dto.LinkDTO, int64, bool, error) {

	links, total, err := lk.repo.GetAllLinks(ctx, page, limit, startDate, endDate)
	if err != nil {
		return nil, 0, true, err
	}

	// Convert sang DTO
	linkDTOs := make([]dto.LinkDTO, len(links))
	for i, l := range links {
		linkDTOs[i] = dto.LinkDTO{
			ID:        l.ID,
			Url:       l.Url,
			Code:      l.Code,
			Link:      l.Link,
			CreatedAt: l.CreatedAt,
		}
	}
	isLast := page*limit >= int(total)

	return linkDTOs, total, isLast, nil
}

func (lk *linkService) DeleteById(ctx context.Context, id uuid.UUID) error {
	err := lk.repo.DeleteById(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
