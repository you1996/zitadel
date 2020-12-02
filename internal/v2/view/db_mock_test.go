package view

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	expectedGetByID                 = `SELECT \* FROM "%s" WHERE \(%s = \$1\) LIMIT 1`
	expectedGetByQuery              = `SELECT \* FROM "%s" WHERE \(LOWER\(%s\) %s LOWER\(\$1\)\) LIMIT 1`
	expectedGetByQueryCaseSensitive = `SELECT \* FROM "%s" WHERE \(%s %s \$1\) LIMIT 1`
	expectedSave                    = `UPDATE "%s" SET "test" = \$1 WHERE "%s"."%s" = \$2`
	expectedRemove                  = `DELETE FROM "%s" WHERE \(%s = \$1\)`
	expectedRemoveByKeys            = func(i int, table string) string {
		sql := fmt.Sprintf(`DELETE FROM "%s"`, table)
		sql += ` WHERE \(%s = \$1\)`
		for j := 1; j < i; j++ {
			sql = sql + ` AND \(%s = \$` + strconv.Itoa(j+1) + `\)`
		}
		return sql
	}
	expectedRemoveByObject           = `DELETE FROM "%s" WHERE "%s"."%s" = \$1`
	expectedRemoveByObjectMultiplePK = `DELETE FROM "%s" WHERE "%s"."%s" = \$1 AND "%s"."%s" = \$2`
	expectedTruncate                 = `TRUNCATE %s;`
	expectedSearch                   = `SELECT \* FROM "%s" OFFSET 0`
	expectedSearchCount              = `SELECT count\(\*\) FROM "%s"`
	expectedSearchLimit              = `SELECT \* FROM "%s" LIMIT %v OFFSET 0`
	expectedSearchLimitCount         = `SELECT count\(\*\) FROM "%s"`
	expectedSearchOffset             = `SELECT \* FROM "%s" OFFSET %v`
	expectedSearchOffsetCount        = `SELECT count\(\*\) FROM "%s"`
	expectedSearchSorting            = `SELECT \* FROM "%s" ORDER BY %s %s OFFSET 0`
	expectedSearchSortingCount       = `SELECT count\(\*\) FROM "%s"`
	expectedSearchQuery              = `SELECT \* FROM "%s" WHERE \(LOWER\(%s\) %s LOWER\(\$1\)\) OFFSET 0`
	expectedSearchQueryCount         = `SELECT count\(\*\) FROM "%s" WHERE \(LOWER\(%s\) %s LOWER\(\$1\)\)`
	expectedSearchQueryAllParams     = `SELECT \* FROM "%s" WHERE \(LOWER\(%s\) %s LOWER\(\$1\)\) ORDER BY %s %s LIMIT %v OFFSET %v`
	expectedSearchQueryAllParamCount = `SELECT count\(\*\) FROM "%s" WHERE \(LOWER\(%s\) %s LOWER\(\$1\)\)`
)

type TestSearchRequest struct {
	limit         uint64
	offset        uint64
	sortingColumn ColumnKey
	asc           bool
	queries       []SearchQuery
}

func (req TestSearchRequest) GetLimit() uint64 {
	return req.limit
}

func (req TestSearchRequest) GetOffset() uint64 {
	return req.offset
}

func (req TestSearchRequest) GetSortingColumn() ColumnKey {
	return req.sortingColumn
}

func (req TestSearchRequest) GetAsc() bool {
	return req.asc
}

func (req TestSearchRequest) GetQueries() []SearchQuery {
	return req.queries
}

type TestSearchQuery struct {
	key    TestSearchKey
	method SearchMethod
	value  string
}

func (req TestSearchQuery) GetKey() ColumnKey {
	return req.key
}

func (req TestSearchQuery) GetMethod() SearchMethod {
	return req.method
}

func (req TestSearchQuery) GetValue() interface{} {
	return req.value
}

type TestSearchKey int32

const (
	TestSearchKey_UNDEFINED TestSearchKey = iota
	TestSearchKey_TEST
	TestSearchKey_ID
)

func (key TestSearchKey) ToColumnName() string {
	switch TestSearchKey(key) {
	case TestSearchKey_TEST:
		return "test"
	case TestSearchKey_ID:
		return "id"
	default:
		return ""
	}
}

type Test struct {
	ID   string `json:"-" gorm:"column:primary_id;primary_key"`
	Test string `json:"test" gorm:"column:test"`
}

type TestMultiplePK struct {
	TestID  string `gorm:"column:testId;primary_key"`
	HodorID string `gorm:"column:hodorId;primary_key"`
	Test    string `gorm:"column:test"`
}

type dbMock struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (db *dbMock) close() {
	db.db.Close()
}

func mockDB(t *testing.T) *dbMock {
	mockDB := dbMock{}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error occured while creating stub db %v", err)
	}

	mockDB.mock = mock
	mockDB.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error occured while connecting to stub db: %v", err)
	}

	mockDB.mock.MatchExpectationsInOrder(true)

	return &mockDB
}

func (db *dbMock) expectGetSearchRequestNoParams(table string, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearch, table)
	queryCount := fmt.Sprintf(expectedSearchCount, table)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestWithLimit(table string, limit, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearchLimit, table, limit)
	queryCount := fmt.Sprintf(expectedSearchLimitCount, table)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestWithOffset(table string, offset, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearchOffset, table, offset)
	queryCount := fmt.Sprintf(expectedSearchOffsetCount, table)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestWithSorting(table, sorting string, sortingColumn ColumnKey, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearchSorting, table, sortingColumn.ToColumnName(), sorting)
	queryCount := fmt.Sprintf(expectedSearchSortingCount, table)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestWithSearchQuery(table, key, method, value string, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearchQuery, table, key, method)
	queryCount := fmt.Sprintf(expectedSearchQueryCount, table, key, method)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WithArgs(value).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WithArgs(value).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestWithAllParams(table, key, method, value, sorting string, sortingColumn ColumnKey, limit, offset, resultAmount, total int) *dbMock {
	query := fmt.Sprintf(expectedSearchQueryAllParams, table, key, method, sortingColumn.ToColumnName(), sorting, limit, offset)
	queryCount := fmt.Sprintf(expectedSearchQueryAllParamCount, table, key, method)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WithArgs(value).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WithArgs(value).
		WillReturnRows(rows)
	return db
}

func (db *dbMock) expectGetSearchRequestErr(table string, resultAmount, total int, err error) *dbMock {
	query := fmt.Sprintf(expectedSearch, table)
	queryCount := fmt.Sprintf(expectedSearchCount, table)

	rows := sqlmock.NewRows([]string{"id"})
	for i := 0; i < resultAmount; i++ {
		rows.AddRow(fmt.Sprintf("hodor-%d", i))
	}

	db.mock.ExpectQuery(queryCount).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))
	db.mock.ExpectQuery(query).
		WillReturnError(err)
	return db
}
