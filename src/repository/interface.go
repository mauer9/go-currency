package repository

type MainStore interface {
	Mssql() MssqlRepository
}
