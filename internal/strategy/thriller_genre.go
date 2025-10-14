package strategy

import (
	"context"
	"github.com/HowAboutAQuiz/go-examples/internal/enums"
)

type ThrillerGenre struct {
}

func NewThrillerGenre() GenreHandler {
	return ThrillerGenre{}
}

func (ag ThrillerGenre) GetGenre() enums.Genre {
	return enums.THRILLER
}

func (ag ThrillerGenre) Handle(ctx context.Context) string {
	return "Thriller"
}
