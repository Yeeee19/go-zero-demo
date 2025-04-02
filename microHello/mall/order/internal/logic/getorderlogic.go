package logic

import (
	"context"
	"user/types/user"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderRequest) (resp *types.OrderResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.getOrderById(req.Id)
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.OrderResponse{
		Id:       req.Id,
		Name:     "999",
		UserName: userResponse.Name,
	}, nil
}

func (l *GetOrderLogic) getOrderById(id string) string {
	return "1"
}
