package controllers

import (
	"encoding/json"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type (
	GenreController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}

	GenreControllerImpl struct {
		genreService service.GenreService
	}
)

func NewGenreController(genreService service.GenreService) GenreController {
	return GenreControllerImpl{
		genreService: genreService,
	}
}

func (gc GenreControllerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	var genre = enums.UNSPECIFIED
	queryGenre := r.URL.Query().Get("genre")
	if strings.TrimSpace(queryGenre) != "" {
		genreEnum, err := strconv.Atoi(queryGenre)
		if err == nil {
			genre = enums.Genre(genreEnum)
		}
	}

	res := gc.genreService.HandleGenre(r.Context(), genre)
	type result struct {
		Message string `json:"message"`
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&result{
		Message: res,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
