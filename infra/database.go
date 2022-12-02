package infra

import "database/sql"

type Database struct {
	dbURI    string
	dbEngine string
}

func NewDatabase(dbURI, dbEngine string) *Database {
	return &Database{
		dbURI:    dbURI,
		dbEngine: dbEngine,
	}
}

func (db *Database) GetDBConnection() (dbConn *sql.DB, err error) {
	dbConn, err = sql.Open(db.dbEngine, db.dbURI)
	if err != nil {
		panic(err)
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)

	return
}
