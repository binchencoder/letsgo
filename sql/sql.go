package sql

import (
	"context"
	"database/sql"

	"upper.io/db.v3/lib/sqlbuilder"
)

// DBSession defines the virtual interface for a database session.
type DBSession interface {
	// Ping returns an error if the database manager could be reached.
	Ping() error

	// PingContext verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	PingContext(ctx context.Context) error

	// Close closes the currently active connection to the database
	// and clears caches.
	Close() error
}

// DBSessionToRawDB converts database session s to the raw *sql.DB.
func DBSessionToRawDB(s DBSession) *sql.DB {
	if db, ok := s.(*sql.DB); ok {
		return db
	}
	return nil
}

// DBSessionToUpperIoDB converts database session s to the upper.io
// sqlbuilder.Database.
func DBSessionToUpperIoDB(s DBSession) sqlbuilder.Database {
	if db, ok := s.(sqlbuilder.Database); ok {
		return db
	}
	return nil
}
