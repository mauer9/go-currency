package mssql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"test/src/repository"
	"test/src/utils/errors"
)

// Main store entity
type Store struct {
	db     *sql.DB
	logger log.Logger

	MssqlRepository repository.MssqlRepository
}

// Store constructor
func NewStore(ctx context.Context, db *sql.DB, logger log.Logger) (*Store, error) {
	repo := &Store{
		logger: logger,
		db:     db,
	}

	err := repo.migrate(ctx)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *Store) migrate(ctx context.Context) error {
	conn, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.InternalServerError.SetDevMessage(err.Error())
	}

	for i := 0; i < len(migrations); i++ {
		_, err = conn.ExecContext(ctx, migrations[i])
		if err != nil {
			_ = conn.Rollback()
			return errors.DBWriteError.SetDevMessage(err.Error())
		}
	}

	_ = conn.Commit()

	return nil
}

func (s *Store) Mssql() repository.MssqlRepository {
	if s.MssqlRepository != nil {
		return s.MssqlRepository
	}

	s.MssqlRepository = &CurrencyRepo{
		store: s,
	}

	return s.MssqlRepository
}
