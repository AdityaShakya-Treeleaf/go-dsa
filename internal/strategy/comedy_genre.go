package strategy

import (
	"context"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
)

type ComedyGenre struct {
}

func NewComedyGenre() GenreHandler {
	return ComedyGenre{}
}

func (ag ComedyGenre) GetGenre() enums.Genre {
	return enums.COMEDY
}

func (ag ComedyGenre) Handle(ctx context.Context) string {
	return "Comedy"
}
