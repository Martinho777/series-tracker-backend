package models

type Series struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Genero      string `json:"genero"`
	Anio        int    `json:"anio"`
	Temporadas  int    `json:"temporadas"`
	ImagenURL   string `json:"imagen_url"`
	Descripcion string `json:"descripcion"`
}