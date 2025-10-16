package strategy

import (
	"context"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
)

type BiographyGenre struct {
}

func NewBiographyGenre() GenreHandler {
	return BiographyGenre{}
}

func (ag BiographyGenre) GetGenre() enums.Genre {
	return enums.BIOGRAPHY
}

func (ag BiographyGenre) Handle(ctx context.Context) string {
	return "Biography"
}
