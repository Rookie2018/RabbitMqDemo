package logic

import (
	"context"
	"encoding/json"

	"rabbitmq-demo/test/internal/svc"
	"rabbitmq-demo/test/pb"

	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestQueueLogic {
	return &TestQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建消息队列
func (l *TestQueueLogic) TestQueue(in *pb.TestReq) (*pb.TestResp, error) {
	// todo: add your logic here and delete this line
	var (
		exchange   = "test_exchange"
		routingKey = "test_normal"
	)
	msgJson, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.RabbitMq.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			Body:        msgJson,
		},
	)

	if err != nil {
		return nil, err
	}

	return &pb.TestResp{
		Ack: 1,
	}, nil
}
