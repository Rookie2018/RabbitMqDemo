package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"rabbitmq-demo/test/internal/svc"
	"rabbitmq-demo/test/pb"

	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestDelayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestDelayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestDelayLogic {
	return &TestDelayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 延迟队列
func (l *TestDelayLogic) TestDelay(in *pb.TestDelayReq) (*pb.TestDelayResp, error) {
	// rabbit.InitDelay("test_delay_queue", "test_delay_exchange", "test_delay")

	fmt.Println("当前数据:", in)

	var (
		exchange   = "test_delay_exchange"
		routingKey = "test_delay"
	)

	msgJson, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	headers := make(amqp.Table)
	headers["x-delay"] = 5 * 60 * 1000

	err = l.svcCtx.RabbitMq.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			Body:        msgJson,
			Headers:     headers,
		},
	)

	if err != nil {
		return nil, err
	}

	return &pb.TestDelayResp{
		Ack: 1,
	}, nil
}
