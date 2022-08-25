package trip

import (
	"context"
	"coolcar/shared/auth"
	"coolcar/shared/id"
)

type impersonation struct {
	AccountID id.AccountID
}

func (i *impersonation) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		auth.ImpersonateAccountHeader: i.AccountID.String(),
	}, nil
}

func (i *impersonation) RequireTransportSecurity() bool {
	return false
}
