package routes

import (
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/controllers"
	"net/http"
)

type GenreRoute struct {
	controller controllers.GenreController
}

func NewGenreController(controller controllers.GenreController) GenreRoute {
	return GenreRoute{
		controller: controller,
	}
}

func (gr GenreRoute) GenreRoute(mux *http.ServeMux) {
	mux.HandleFunc("/run", gr.controller.Handle)
}
