package service

import (
	"context"
	"github.com/HowAboutAQuiz/go-examples/internal/enums"
	"github.com/HowAboutAQuiz/go-examples/internal/strategy"
	"log"
)

type (
	GenreService interface {
		HandleGenre(ctx context.Context, genre enums.Genre) string
	}

	GenreServiceImpl struct {
		genreFactory strategy.GenreHandlerFactory
	}
)

func NewGenreService(genreFactory strategy.GenreHandlerFactory) GenreService {
	return GenreServiceImpl{
		genreFactory: genreFactory,
	}
}

func (gs GenreServiceImpl) HandleGenre(ctx context.Context, genre enums.Genre) string {
	log.Default().Printf("Handling for genre: %d", genre)
	handler := gs.genreFactory.GetGenreHandler(genre)
	return handler.Handle(ctx)
}
