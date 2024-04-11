package mssql

import (
	"context"
	"database/sql"
	"errors"
	"test/src/domain"
)

type CurrencyRepo struct {
	store *Store
}

func (c *CurrencyRepo) UpdateCurrencyByDate(ctx context.Context, currencies []domain.Currency) {
	tx, err := c.store.db.BeginTx(ctx, nil)
	if err != nil {
		_ = c.store.logger.Log("err", err.Error())
	}

	mergeStmt := `
		MERGE INTO R_CURRENCY AS target
		USING (
			SELECT @TTITLE AS TITLE, @CODE AS CODE, @A_DATE AS A_DATE, @VALUE AS VALUE
		) AS source
		ON target.A_DATE = source.A_DATE
		WHEN MATCHED THEN
			UPDATE SET
				target.TITLE = source.TITLE,
				target.CODE = source.CODE,
			    target.VALUE = source.VALUE,
			    target.A_DATE = source.A_DATE
		WHEN NOT MATCHED BY TARGET THEN
			INSERT (TITLE, CODE, VALUE, A_DATE)
			VALUES (source.TITLE, source.CODE, source.VALUE, source.A_DATE);
	`

	for _, cr := range currencies {
		_, err = tx.Exec(mergeStmt,
			sql.Named("TITLE", cr.Title),
			sql.Named("CODE", cr.Code),
			sql.Named("A_DATE", cr.ADate),
			sql.Named("VALUE", cr.Value),
		)

		if err != nil {
			_ = c.store.logger.Log("err", err.Error())
		}
	}
}

func (c *CurrencyRepo) GetCurrenciesByDateAndCode(ctx context.Context, date, code string) ([]domain.Currency, error) {
	tx, err := c.store.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, errors.New("упс, внутренняя ошибка")
	}

	var resp []domain.Currency

	rows, err := tx.QueryContext(ctx, "SELECT * FROM R_CURRENCY WHERE A_DATE = ? AND CODE = ?", date, code)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.New("упс, внутренняя ошибка")
	}

	for rows.Next() {
		var cr domain.Currency
		err = rows.Scan(
			&cr.ID,
			&cr.Title,
			&cr.Code,
			&cr.Value,
			&cr.ADate,
		)

		if err != nil {
			_ = tx.Rollback()
			return nil, errors.New("что то пошло не так")
		}

		resp = append(resp, cr)
	}

	_ = tx.Commit()

	return resp, nil
}
