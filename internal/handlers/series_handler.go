package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"series-tracker-backend/internal/service"
	"series-tracker-backend/internal/utils"
)

type SeriesHandler struct {
	Service *service.SeriesService
}

func NewSeriesHandler(service *service.SeriesService) *SeriesHandler {
	return &SeriesHandler{Service: service}
}

func (h *SeriesHandler) GetAllSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	seriesList, err := h.Service.GetAllSeries()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "No se pudieron obtener las series")
		return
	}

	utils.WriteJSON(w, http.StatusOK, seriesList)
}

func (h *SeriesHandler) GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) != 2 || pathParts[0] != "series" {
		utils.WriteError(w, http.StatusBadRequest, "Ruta inválida")
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "El id debe ser un número válido")
		return
	}

	serie, err := h.Service.GetSeriesByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "No se pudo obtener la serie")
		return
	}

	if serie == nil {
		utils.WriteError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	utils.WriteJSON(w, http.StatusOK, serie)
}