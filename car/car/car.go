package car

import (
	"coolcar/car/dao"
	"coolcar/car/mq"
	"go.uber.org/zap"
)

type Service struct {
	Logger    *zap.Logger
	Mongo     *dao.Mongo
	Publisher mq.Publisher
}
