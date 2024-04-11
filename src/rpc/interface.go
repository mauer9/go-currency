package rpc

import (
	"context"
	"test/src/domain"
)

type RPC interface {
	GetCurrencyRate(_ context.Context, date string) ([]domain.Currency, error)
}
