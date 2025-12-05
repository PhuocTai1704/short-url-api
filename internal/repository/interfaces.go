package repository

import model "short-url-api/internal/models"

type LinkRepo interface {
	Create(link *model.Link) error
	
	FirstOrCreate(link *model.Link) error
}
