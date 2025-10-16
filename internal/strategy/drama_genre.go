package strategy

import (
	"context"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
)

type DramaGenre struct {
}

func NewDramaGenre() GenreHandler {
	return DramaGenre{}
}

func (ag DramaGenre) GetGenre() enums.Genre {
	return enums.DRAMA
}

func (ag DramaGenre) Handle(ctx context.Context) string {
	return "Drama"
}
