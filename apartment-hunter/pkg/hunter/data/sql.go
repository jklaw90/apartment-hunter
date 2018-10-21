package data

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type SQL struct {
	writer *sql.DB
	reader *sql.DB
}

func NewSQL(rURI, wURI string) *SQL {
	tries := 0
	var writeDB *sql.DB
	var err error
	for {
		if writeDB, err = sql.Open("postgres", wURI); err == nil {
			err = writeDB.Ping()
		}
		if err != nil {
			tries += 1
			if tries > 20 {
				panic(err)
			}
			time.Sleep(time.Duration(5) * time.Second)
		} else {
			break
		}
	}
	return &SQL{
		writer: writeDB,
		reader: writeDB,
	}
}

func (s *SQL) Reader() *sql.DB {
	return s.reader
}

func (s *SQL) Writer() *sql.DB {
	return s.writer
}

func (s *SQL) Ping() error {
	err := s.reader.Ping()
	if err != nil {
		return err
	}
	err = s.writer.Ping()
	return err
}

func Prepare(stmt string, db DB) *sql.Stmt {
	resp, err := db.Reader().Prepare(stmt)
	if err != nil {
		panic(err)
	}
	return resp
}
