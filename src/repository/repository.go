package repository

import (
	"context"
	"test/src/domain"
)

type MssqlRepository interface {
	UpdateCurrencyByDate(ctx context.Context, currencies []domain.Currency)
	GetCurrenciesByDateAndCode(ctx context.Context, date, code string) ([]domain.Currency, error)
}
