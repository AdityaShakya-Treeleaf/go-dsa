package main

import (
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/controllers"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/dsa"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/routes"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/service"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/strategy"
	"net/http"
)

func main() {

	adjacentIncreasingSubArraysDetection := dsa.AdjacentIncreasingSubArraysDetection{}
	adjacentIncreasingSubArraysDetectionV2 := dsa.AdjacentIncreasingSubArraysDetectionV2{}
	adjacentIncreasingSubArraysDetection.Execute()
	adjacentIncreasingSubArraysDetectionV2.Execute()
	genreStrategy := strategy.NewGenreHandlerFactory()
	genreStrategy.AddGenreHandler(strategy.NewActionGenre())
	genreStrategy.AddGenreHandler(strategy.NewComedyGenre())
	genreStrategy.AddGenreHandler(strategy.NewThrillerGenre())
	genreStrategy.AddGenreHandler(strategy.NewHorrorGenre())
	genreStrategy.AddGenreHandler(strategy.NewDramaGenre())
	genreStrategy.AddGenreHandler(strategy.NewBiographyGenre())

	genreService := service.NewGenreService(genreStrategy)
	genreController := controllers.NewGenreController(genreService)
	genreRoute := routes.NewGenreController(genreController)

	mux := http.NewServeMux()
	genreRoute.GenreRoute(mux)
	//
	//log.Default().Printf("Listening at 8081...")
	//log.Fatal(http.ListenAndServe(":8081", mux))
}
