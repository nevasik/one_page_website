package _interface

import (
	"Odnostranishka/internal/plugin/model"
	"context"
)

type UC interface {
	Create(ctx context.Context, req model.Order) (model.OrderResp, int)
}
