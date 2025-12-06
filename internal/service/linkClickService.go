package service

import (
	"context"
	"net/http"
	model "short-url-api/internal/models"
	"short-url-api/internal/repository"
	"time"

	"github.com/mssola/user_agent"

	"github.com/google/uuid"
)

type linkClickService struct {
	repo repository.LinkClickRepo
}

func NewLinkClickService(repo repository.LinkClickRepo) LinkClickService {
	return &linkClickService{
		repo: repo,
	}
}

func (lc *linkClickService) IncreaseClick(ctx context.Context, linkId uuid.UUID, req *http.Request) error {
	// Lấy IP
	ip := req.RemoteAddr
	// Nếu bạn sử dụng proxy/nginx, có thể lấy từ X-Forwarded-For
	if forwarded := req.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = forwarded
	}

	// Lấy Referrer
	referrer := req.Referer()

	// Lấy ngôn ngữ trình duyệt
	language := req.Header.Get("Accept-Language")

	// Phân tích User-Agent
	device, os, browser := parseUserAgent(req.UserAgent())

	click := &model.LinkClick{
		LinkID:    linkId,
		ClickedAt: time.Now(),
		IP:        ip,
		Referrer:  referrer,
		Country:   "", // Nếu muốn lấy, tích hợp GeoIP
		City:      "", // Nếu muốn lấy, tích hợp GeoIP
		Device:    device,
		Language:  language,
		OS:        os,
		Browser:   browser,
		CreatedAt: time.Now(),
	}

	if err := lc.repo.Create(ctx, click); err != nil {
		return err
	}

	return nil
}

// parseUserAgent giữ nguyên
func parseUserAgent(uaString string) (device, os, browser string) {
	ua := user_agent.New(uaString)
	browserName, _ := ua.Browser()
	os = ua.OS()

	if ua.Mobile() {
		device = "mobile"
	} else if ua.Platform() == "iPad" || ua.Platform() == "Tablet" {
		device = "tablet"
	} else {
		device = "desktop"
	}

	return device, os, browserName
}
