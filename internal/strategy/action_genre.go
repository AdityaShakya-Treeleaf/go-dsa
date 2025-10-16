package strategy

import (
	"context"
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/enums"
)

type ActionGenre struct {
}

func NewActionGenre() GenreHandler {
	return ActionGenre{}
}

func (ag ActionGenre) GetGenre() enums.Genre {
	return enums.ACTION
}

func (ag ActionGenre) Handle(ctx context.Context) string {
	return "Action"
}
