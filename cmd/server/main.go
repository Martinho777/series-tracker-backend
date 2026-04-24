package main

import (
	"fmt"
	"net/http"

	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/handlers"
	"series-tracker-backend/internal/middleware"
	"series-tracker-backend/internal/repository"
	"series-tracker-backend/internal/service"
)

func main() {
	dbConn, err := db.ConnectPostgres()
	if err != nil {
		fmt.Println("Error de conexión a la base de datos:", err)
		return
	}
	defer dbConn.Close()

	seriesRepo := repository.NewSeriesRepository(dbConn)
	seriesService := service.NewSeriesService(seriesRepo)
	seriesHandler := handlers.NewSeriesHandler(seriesService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message":"API del Series Tracker activa"}`)
	})

	mux.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			seriesHandler.GetAllSeries(w, r)
		case http.MethodPost:
			seriesHandler.CreateSeries(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, `{"error":"Método no permitido"}`)
		}
	})

	mux.HandleFunc("/series/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			seriesHandler.GetSeriesByID(w, r)
		case http.MethodPut:
			seriesHandler.UpdateSeries(w, r)
		case http.MethodDelete:
			seriesHandler.DeleteSeries(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, `{"error":"Método no permitido"}`)
		}
	})

	mux.HandleFunc("/upload", seriesHandler.UploadImage)
	
	mux.Handle("/openapi.yaml", http.FileServer(http.Dir(".")))
	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})

	handlerWithCORS := middleware.EnableCORS(mux)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	err = http.ListenAndServe(":8080", handlerWithCORS)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}