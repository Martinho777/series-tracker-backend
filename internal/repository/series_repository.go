package repository

import (
	"database/sql"
	"fmt"

	"series-tracker-backend/internal/models"
)

type SeriesRepository struct {
	DB *sql.DB
}

func NewSeriesRepository(db *sql.DB) *SeriesRepository {
	return &SeriesRepository{DB: db}
}

func (r *SeriesRepository) GetAll() ([]models.Series, error) {
	query := `
		SELECT id, titulo, genero, anio, temporadas, imagen_url, descripcion
		FROM series
		ORDER BY id ASC
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar las series: %w", err)
	}
	defer rows.Close()

	var seriesList []models.Series

	for rows.Next() {
		var serie models.Series

		err := rows.Scan(
			&serie.ID,
			&serie.Titulo,
			&serie.Genero,
			&serie.Anio,
			&serie.Temporadas,
			&serie.ImagenURL,
			&serie.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("error al leer una serie: %w", err)
		}

		seriesList = append(seriesList, serie)
	}

	return seriesList, nil
}

func (r *SeriesRepository) GetByID(id int) (*models.Series, error) {
	query := `
		SELECT id, titulo, genero, anio, temporadas, imagen_url, descripcion
		FROM series
		WHERE id = $1
	`

	var serie models.Series

	err := r.DB.QueryRow(query, id).Scan(
		&serie.ID,
		&serie.Titulo,
		&serie.Genero,
		&serie.Anio,
		&serie.Temporadas,
		&serie.ImagenURL,
		&serie.Descripcion,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al consultar la serie por id: %w", err)
	}

	return &serie, nil
}

func (r *SeriesRepository) Create(serie *models.Series) (*models.Series, error) {
	query := `
		INSERT INTO series (titulo, genero, anio, temporadas, imagen_url, descripcion)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, titulo, genero, anio, temporadas, imagen_url, descripcion
	`

	var createdSeries models.Series

	err := r.DB.QueryRow(
		query,
		serie.Titulo,
		serie.Genero,
		serie.Anio,
		serie.Temporadas,
		serie.ImagenURL,
		serie.Descripcion,
	).Scan(
		&createdSeries.ID,
		&createdSeries.Titulo,
		&createdSeries.Genero,
		&createdSeries.Anio,
		&createdSeries.Temporadas,
		&createdSeries.ImagenURL,
		&createdSeries.Descripcion,
	)

	if err != nil {
		return nil, fmt.Errorf("error al crear la serie: %w", err)
	}

	return &createdSeries, nil
}

func (r *SeriesRepository) Update(id int, serie *models.Series) (*models.Series, error) {
	query := `
		UPDATE series
		SET titulo = $1,
			genero = $2,
			anio = $3,
			temporadas = $4,
			imagen_url = $5,
			descripcion = $6,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
		RETURNING id, titulo, genero, anio, temporadas, imagen_url, descripcion
	`

	var updatedSeries models.Series

	err := r.DB.QueryRow(
		query,
		serie.Titulo,
		serie.Genero,
		serie.Anio,
		serie.Temporadas,
		serie.ImagenURL,
		serie.Descripcion,
		id,
	).Scan(
		&updatedSeries.ID,
		&updatedSeries.Titulo,
		&updatedSeries.Genero,
		&updatedSeries.Anio,
		&updatedSeries.Temporadas,
		&updatedSeries.ImagenURL,
		&updatedSeries.Descripcion,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al actualizar la serie: %w", err)
	}

	return &updatedSeries, nil
}