package sim

import (
	"context"
	coolenvpb "coolcar/shared/coolenv"
)

type PosSubscriber interface {
	Subscribe(context.Context) (ch chan *coolenvpb.CarPosUpdate, cleanUp func(), err error)
}
