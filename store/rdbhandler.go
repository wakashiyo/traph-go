package store

import (
	"database/sql"

	"github.com/wakashiyo/traph-go/interfaces"
)

type Handler struct {
	cnn *sql.DB
}

func NewDB(driver string, source string) (interfaces.Repository, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	handler := new(Handler)
	handler.cnn = db
	return handler, nil
}

func (h *Handler) Execute(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := h.cnn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	sqlresult := new(SqlResult)
	sqlresult.Result = result
	return sqlresult, nil
}

func (h *Handler) Query(statement string) (interfaces.Rows, error) {
	rows, err := h.cnn.Query(statement)
	if err != nil {
		return nil, err
	}
	row := new(SqlRows)
	row.Rows = rows
	return row, nil
}

func (h *Handler) QueryRow(statement string, args ...interface{}) interfaces.Row {
	row := h.cnn.QueryRow(statement, args...)
	_row := new(SqlRow)
	_row.Row = row
	return _row
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRows struct {
	Rows *sql.Rows
}

func (r SqlRows) Scan(args ...interface{}) error {
	return r.Rows.Scan(args...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

type SqlRow struct {
	Row *sql.Row
}

func (r SqlRow) Scan(args ...interface{}) error {
	return r.Row.Scan(args...)
}
