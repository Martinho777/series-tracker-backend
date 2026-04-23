package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"series-tracker-backend/internal/models"
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

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page == 0 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	filters := models.SeriesFilters{
		Query: r.URL.Query().Get("q"),
		Sort:  r.URL.Query().Get("sort"),
		Order: r.URL.Query().Get("order"),
		Page:  page,
		Limit: limit,
	}

	seriesList, err := h.Service.GetAllSeries(filters)
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

func (h *SeriesHandler) CreateSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var serie models.Series

	err := json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "El cuerpo de la solicitud debe ser JSON válido")
		return
	}

	createdSeries, err := h.Service.CreateSeries(&serie)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdSeries)
}

func (h *SeriesHandler) UpdateSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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

	var serie models.Series

	err = json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "El cuerpo de la solicitud debe ser JSON válido")
		return
	}

	updatedSeries, err := h.Service.UpdateSeries(id, &serie)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if updatedSeries == nil {
		utils.WriteError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedSeries)
}

func (h *SeriesHandler) DeleteSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	deleted, err := h.Service.DeleteSeries(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !deleted {
		utils.WriteError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}