package interfaces

import (
	"testing"
)

type RepositoryStub struct {}

func (r *RepositoryStub) Execute(query string, args ...interface{}) (Result, error) {
	result := new(RepositoryStub)
	result.resultId = 1
	return result, nil
}

func (r *RepositoryStub) Query(statement string) (Rows, error) {
	rows := new(RowsStub)
	return rows, nil
}

func (r *RepositoryStub) QueryRow(statement string, args ...interface{}) Row {
	row := new(RowStub)
	return row, nil
}

type ResultStub struct {
	resultId: int64
}

func (r *ResultStub) LastInsertId() (int64, error) {
	return r.resultId, nil
}

func (r *ResultStub) RowsAffected() (int64, error) {
	return r.resultId, nil
}

type RowsStub struct {}

func (r *RowsStub) Scan(args ...interface{}) error {
	return nil
}

func (r *RowsStub) Next() bool {
	return true
}

func (r *RowsStub) Close() error {
	return true
}

type RowStub struct {}

func (r *RowStub) Scan(args ...interface{}) error {
	return nil
}

func TestUserStore(t *testing.T) {
	repoImpl := ReposirotyImpl{
		repository: RepositoryStub{}
	}
	//
}

func TestUserUpdate(t *testing.T) {
	repoImpl := ReposirotyImpl{
		repository: RepositoryStub{}
	}
	//
}
