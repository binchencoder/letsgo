package sql

import (
	"context"
	"database/sql"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	updb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// mockUpperIoDB implements DBSession, but doesn't implement either sql.DB
// or sqlbuilder.Database.
type mockInvalidDB struct{}

func (midb *mockInvalidDB) Ping() error {
	return nil
}

func (midb *mockInvalidDB) PingContext(ctx context.Context) error {
	return nil
}

func (midb *mockInvalidDB) Close() error {
	return nil
}

// mockUpperIoDB implements interface sqlbuilder.Database.
type mockUpperIoDB struct{}

func (muidb *mockUpperIoDB) Ping() error {
	return nil
}

func (muidb *mockUpperIoDB) PingContext(ctx context.Context) error {
	return nil
}

func (muidb *mockUpperIoDB) Close() error {
	return nil
}

func (muidb *mockUpperIoDB) Driver() interface{} {
	return nil
}

func (muidb *mockUpperIoDB) Open(updb.ConnectionURL) error {
	return nil
}

func (muidb *mockUpperIoDB) Collection(string) updb.Collection {
	return nil
}

func (muidb *mockUpperIoDB) Collections() ([]string, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) Name() string {
	return ""
}

func (muidb *mockUpperIoDB) ConnectionURL() updb.ConnectionURL {
	return nil
}

func (muidb *mockUpperIoDB) ClearCache() {}

func (muidb *mockUpperIoDB) SetLogging(bool) {}

func (muidb *mockUpperIoDB) LoggingEnabled() bool {
	return false
}

func (muidb *mockUpperIoDB) SetLogger(updb.Logger) {}

func (muidb *mockUpperIoDB) Logger() updb.Logger {
	return nil
}

func (muidb *mockUpperIoDB) SetPreparedStatementCache(bool) {}

func (muidb *mockUpperIoDB) PreparedStatementCacheEnabled() bool {
	return false
}

func (muidb *mockUpperIoDB) SetConnMaxLifetime(time.Duration) {}

func (muidb *mockUpperIoDB) ConnMaxLifetime() time.Duration {
	return 0
}

func (muidb *mockUpperIoDB) SetMaxIdleConns(int) {}

func (muidb *mockUpperIoDB) MaxIdleConns() int {
	return 0
}

func (muidb *mockUpperIoDB) SetMaxOpenConns(int) {}

func (muidb *mockUpperIoDB) MaxOpenConns() int {
	return 0
}

func (muidb *mockUpperIoDB) Select(columns ...interface{}) sqlbuilder.Selector {
	return nil
}

func (muidb *mockUpperIoDB) SelectFrom(table ...interface{}) sqlbuilder.Selector {
	return nil
}

func (muidb *mockUpperIoDB) InsertInto(table string) sqlbuilder.Inserter {
	return nil
}

func (muidb *mockUpperIoDB) DeleteFrom(table string) sqlbuilder.Deleter {
	return nil
}

func (muidb *mockUpperIoDB) Update(table string) sqlbuilder.Updater {
	return nil
}

func (muidb *mockUpperIoDB) Exec(query interface{}, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) ExecContext(ctx context.Context, query interface{}, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) Prepare(query interface{}) (*sql.Stmt, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) PrepareContext(ctx context.Context, query interface{}) (*sql.Stmt, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) Query(query interface{}, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) QueryContext(ctx context.Context, query interface{}, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) QueryRow(query interface{}, args ...interface{}) (*sql.Row, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) QueryRowContext(ctx context.Context, query interface{}, args ...interface{}) (*sql.Row, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) Iterator(query interface{}, args ...interface{}) sqlbuilder.Iterator {
	return nil
}

func (muidb *mockUpperIoDB) IteratorContext(ctx context.Context, query interface{}, args ...interface{}) sqlbuilder.Iterator {
	return nil
}

func (muidb *mockUpperIoDB) NewTx(ctx context.Context) (sqlbuilder.Tx, error) {
	return nil, nil
}

func (muidb *mockUpperIoDB) Tx(ctx context.Context, fn func(sess sqlbuilder.Tx) error) error {
	return nil
}

func (muidb *mockUpperIoDB) SetTxOptions(sql.TxOptions) {}

func (muidb *mockUpperIoDB) Context() context.Context {
	return nil
}

func (muidb *mockUpperIoDB) WithContext(context.Context) sqlbuilder.Database {
	return nil
}

func (muidb *mockUpperIoDB) TxOptions() *sql.TxOptions {
	return nil
}

func TestDBSessionToRawDB(t *testing.T) {
	Convey("Convert DBSession to raw *sql.DB", t, func() {
		Convey("when the session is a valid *sql.DB instance", func() {
			s := sql.DB{}
			db := DBSessionToRawDB(&s)
			So(db, ShouldNotBeNil)
		})
		Convey("when the session is not a valid *sql.DB instance", func() {
			s := mockInvalidDB{}
			db := DBSessionToRawDB(&s)
			So(db, ShouldBeNil)
		})
	})
}

// func TestDBSessionToUpperIoDB(t *testing.T) {
// 	Convey("Convert DBSession to upper.io sqlbuilder.Database", t, func() {
// 		Convey("when the session is a valid sqlbuilder.Database instance", func() {
// 			var s sqlbuilder.Database = &mockUpperIoDB{}
// 			db := DBSessionToUpperIoDB(s)
// 			So(db, ShouldNotBeNil)
// 		})
// 		Convey("when the session is not a valid sqlbuilder.Database instance", func() {
// 			s := mockInvalidDB{}
// 			db := DBSessionToUpperIoDB(&s)
// 			So(db, ShouldBeNil)
// 		})
// 	})
// }
