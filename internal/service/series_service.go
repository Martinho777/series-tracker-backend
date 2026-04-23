package service

import (
	"fmt"
	"strings"

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

func (s *SeriesService) CreateSeries(serie *models.Series) (*models.Series, error) {
	serie.Titulo = strings.TrimSpace(serie.Titulo)
	serie.Genero = strings.TrimSpace(serie.Genero)
	serie.Descripcion = strings.TrimSpace(serie.Descripcion)
	serie.ImagenURL = strings.TrimSpace(serie.ImagenURL)

	if serie.Titulo == "" {
		return nil, fmt.Errorf("el título es obligatorio")
	}

	if serie.Genero == "" {
		return nil, fmt.Errorf("el género es obligatorio")
	}

	if serie.Anio < 1900 {
		return nil, fmt.Errorf("el año debe ser mayor o igual a 1900")
	}

	if serie.Temporadas < 1 {
		return nil, fmt.Errorf("las temporadas deben ser al menos 1")
	}

	return s.Repo.Create(serie)
}

func (s *SeriesService) UpdateSeries(id int, serie *models.Series) (*models.Series, error) {
	serie.Titulo = strings.TrimSpace(serie.Titulo)
	serie.Genero = strings.TrimSpace(serie.Genero)
	serie.Descripcion = strings.TrimSpace(serie.Descripcion)
	serie.ImagenURL = strings.TrimSpace(serie.ImagenURL)

	if id < 1 {
		return nil, fmt.Errorf("el id debe ser mayor que 0")
	}

	if serie.Titulo == "" {
		return nil, fmt.Errorf("el título es obligatorio")
	}

	if serie.Genero == "" {
		return nil, fmt.Errorf("el género es obligatorio")
	}

	if serie.Anio < 1900 {
		return nil, fmt.Errorf("el año debe ser mayor o igual a 1900")
	}

	if serie.Temporadas < 1 {
		return nil, fmt.Errorf("las temporadas deben ser al menos 1")
	}

	return s.Repo.Update(id, serie)
}

func (s *SeriesService) DeleteSeries(id int) (bool, error) {
	if id < 1 {
		return false, fmt.Errorf("el id debe ser mayor que 0")
	}

	return s.Repo.Delete(id)
}