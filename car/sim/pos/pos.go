package pos

import (
	"coolcar/car/mq/amqpclt"
	"go.uber.org/zap"
)

type Subscriber struct {
	Sub    *amqpclt.Publisher
	Logger *zap.Logger
}
