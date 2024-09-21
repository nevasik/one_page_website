package uc

import (
	i "Odnostranishka/internal/plugin/interface"
	"Odnostranishka/internal/plugin/model"
	"Odnostranishka/pkg/http/codes"
	"context"
	"fmt"
	lg "gitlab.com/nevasik7/logger"
)

type uc struct{}

func New() i.UC {
	return &uc{}
}

func (u *uc) Create(_ context.Context, req model.Order) (model.OrderResp, int) {
	var (
		resp model.OrderResp
	)

	lg.SendAlert(fmt.Sprintf("Прилетела заявка от %v. \nНомер=%v. \nОписание=%v. \nКонтакты=%v", req.Name, req.Phone, req.Description, req.Telegram), lg.Odnostr)

	return resp, codes.NoError
}
