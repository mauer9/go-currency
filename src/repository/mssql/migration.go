package mssql

var migrations = []string{
	`IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'R_CURRENCY')
		BEGIN
			EXEC('
				CREATE TABLE R_CURRENCY (
					ID int IDENTITY(1,1) PRIMARY KEY,
					TITLE varchar(60) not null,
					CODE varchar(3) not null,
					VALUE numeric(18, 2) not null,
					A_DATE date not null
				)
			');
		END;`,
}
