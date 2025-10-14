package strategy

import (
	"context"
	"github.com/HowAboutAQuiz/go-examples/internal/enums"
	"log"
)

type (
	GenreHandler interface {
		Handle(ctx context.Context) string
		GetGenre() enums.Genre
	}

	GenreHandlerFactory struct {
		genreMap map[enums.Genre]GenreHandler
	}
)

func NewGenreHandlerFactory() GenreHandlerFactory {
	return GenreHandlerFactory{
		genreMap: make(map[enums.Genre]GenreHandler),
	}
}

func (gh GenreHandlerFactory) AddGenreHandler(handler GenreHandler) {
	log.Default().Printf("Adding genre: %d", handler.GetGenre())
	gh.genreMap[handler.GetGenre()] = handler
}

func (gh GenreHandlerFactory) GetGenreHandler(genre enums.Genre) GenreHandler {
	handler, ok := gh.genreMap[genre]
	if !ok {
		panic("No handler found")
	}

	return handler
}
