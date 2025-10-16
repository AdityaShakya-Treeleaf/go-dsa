package strategy

import (
	"context"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
)

type HorrorGenre struct {
}

func NewHorrorGenre() GenreHandler {
	return HorrorGenre{}
}

func (ag HorrorGenre) GetGenre() enums.Genre {
	return enums.HORROR
}

func (ag HorrorGenre) Handle(ctx context.Context) string {
	return "Horror"
}
