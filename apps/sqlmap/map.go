package main

import (
	"solod.dev/so/c"
	"solod.dev/so/errors"
	"solod.dev/so/mem"
	"solod.dev/so/slices"
	"solod.dev/so/strings"
)

var (
	ErrCreate   = errors.New("sqlmap: create schema failed")
	ErrExec     = errors.New("sqlmap: exec failed")
	ErrNotFound = errors.New("sqlmap: not found")
	ErrPrepare  = errors.New("sqlmap: prepare failed")
)

const (
	sqlCreate = "create table if not exists kv (key text primary key, val)"
	sqlGet    = "select val from kv where key = ?"
	sqlSet    = "insert or replace into kv (key, val) values (?, ?)"
	sqlDelete = "delete from kv where key = ?"
)

// SQLMap is a simple key-value store backed by an SQLite database.
type SQLMap struct {
	db *sqlite3
}

type SQLMapResult struct {
	val SQLMap
	err error
}

// NewSQLMap creates a new SQLMap using the provided connection string.
// It opens a connection to the SQLite database and creates the underlying
// key-value table if it does not already exist.
//
// The caller is responsible for calling Close on the returned SQLMap
// when it is no longer needed.
func NewSQLMap(connStr string) (SQLMap, error) {
	var db *sqlite3
	rc := sqlite3_open(connStr, &db)
	if rc != sqliteOK {
		return SQLMap{}, ErrCreate
	}

	rc = sqlite3_exec(db, sqlCreate, nil, nil, nil)
	if rc != sqliteOK {
		sqlite3_close(db)
		return SQLMap{}, ErrCreate
	}
	return SQLMap{db}, nil
}

// GetInt returns the integer value associated with the specified key.
func (m *SQLMap) GetInt(key string) (int, error) {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlGet, -1, &stmt, nil)
	if rc != sqliteOK {
		return 0, ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	rc = sqlite3_step(stmt)
	if rc == sqliteDone {
		return 0, ErrNotFound
	}
	if rc != sqliteRow {
		return 0, ErrExec
	}

	result := int(sqlite3_column_int64(stmt, 0))
	return result, nil
}

// GetFloat64 returns the float64 value associated with the specified key.
func (m *SQLMap) GetFloat64(key string) (float64, error) {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlGet, -1, &stmt, nil)
	if rc != sqliteOK {
		return 0, ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	rc = sqlite3_step(stmt)
	if rc == sqliteDone {
		return 0, ErrNotFound
	}
	if rc != sqliteRow {
		return 0, ErrExec
	}

	result := sqlite3_column_double(stmt, 0)
	return result, nil
}

// GetString returns the string value associated with the specified key.
// The caller owns the returned string and must free it with mem.FreeString.
func (m *SQLMap) GetString(a mem.Allocator, key string) (string, error) {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlGet, -1, &stmt, nil)
	if rc != sqliteOK {
		return "", ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	rc = sqlite3_step(stmt)
	if rc == sqliteDone {
		return "", ErrNotFound
	}
	if rc != sqliteRow {
		return "", ErrExec
	}

	text := c.Val[*c.ConstChar]("(const char*)sqlite3_column_text(stmt, 0)")
	tmp := c.String(text)
	result := strings.Clone(a, tmp)
	return result, nil
}

// GetByte returns the raw blob value associated with the specified key.
// The caller owns the returned slice and must free it with mem.FreeSlice.
func (m *SQLMap) GetByte(a mem.Allocator, key string) ([]byte, error) {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlGet, -1, &stmt, nil)
	if rc != sqliteOK {
		return nil, ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	rc = sqlite3_step(stmt)
	if rc == sqliteDone {
		return nil, ErrNotFound
	}
	if rc != sqliteRow {
		return nil, ErrExec
	}

	ptr := sqlite3_column_blob(stmt, 0).(*byte)
	n := sqlite3_column_bytes(stmt, 0)
	src := c.Bytes(ptr, int(n))
	result := slices.Clone(a, src)
	return result, nil
}

// SetInt stores an integer value for the specified key.
func (m *SQLMap) SetInt(key string, val int) error {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlSet, -1, &stmt, nil)
	if rc != sqliteOK {
		return ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	sqlite3_bind_int64(stmt, 2, int64(val))

	rc = sqlite3_step(stmt)
	if rc != sqliteDone {
		return ErrExec
	}
	return nil
}

// SetFloat64 stores a float64 value for the specified key.
func (m *SQLMap) SetFloat64(key string, val float64) error {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlSet, -1, &stmt, nil)
	if rc != sqliteOK {
		return ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	sqlite3_bind_double(stmt, 2, val)

	rc = sqlite3_step(stmt)
	if rc != sqliteDone {
		return ErrExec
	}
	return nil
}

// SetString stores a string value for the specified key.
func (m *SQLMap) SetString(key string, val string) error {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlSet, -1, &stmt, nil)
	if rc != sqliteOK {
		return ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	sqlite3_bind_text(stmt, 2, val, int32(len(val)), nil)

	rc = sqlite3_step(stmt)
	if rc != sqliteDone {
		return ErrExec
	}
	return nil
}

// SetByte stores a blob value for the specified key.
func (m *SQLMap) SetByte(key string, val []byte) error {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlSet, -1, &stmt, nil)
	if rc != sqliteOK {
		return ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	sqlite3_bind_blob(stmt, 2, val, int32(len(val)), nil)

	rc = sqlite3_step(stmt)
	if rc != sqliteDone {
		return ErrExec
	}
	return nil
}

// Delete removes the entry with the specified key.
func (m *SQLMap) Delete(key string) error {
	var stmt *sqlite3_stmt
	rc := sqlite3_prepare_v2(m.db, sqlDelete, -1, &stmt, nil)
	if rc != sqliteOK {
		return ErrPrepare
	}
	defer sqlite3_finalize(stmt)

	sqlite3_bind_text(stmt, 1, key, int32(len(key)), nil)
	rc = sqlite3_step(stmt)

	if rc != sqliteDone {
		return ErrExec
	}
	return nil
}

// Close releases resources associated with the SQLMap.
func (m *SQLMap) Close() {
	sqlite3_close(m.db)
}
