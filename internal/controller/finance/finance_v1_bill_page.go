package finance

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

func (c *ControllerV1) BillPage(ctx context.Context, req *v1.BillPageReq) (res *v1.BillPageRes, err error) {

	billPageRes, err := service.Finance().BillPage(ctx, req.FinanceBillPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.BillPageRes{
		FinanceBillPageRes: billPageRes,
	}

	return
}
