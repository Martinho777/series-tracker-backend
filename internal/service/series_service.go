package service

import (
	"series-tracker-backend/internal/models"
	"series-tracker-backend/internal/repository"
)

type SeriesService struct {
	Repo *repository.SeriesRepository
}

func NewSeriesService(repo *repository.SeriesRepository) *SeriesService {
	return &SeriesService{Repo: repo}
}

func (s *SeriesService) GetAllSeries() ([]models.Series, error) {
	return s.Repo.GetAll()
}

func (s *SeriesService) GetSeriesByID(id int) (*models.Series, error) {
	return s.Repo.GetByID(id)
}