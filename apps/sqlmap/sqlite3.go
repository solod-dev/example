package main

import "solod.dev/so/c"

//so:include <sqlite3.h>

//so:extern SQLITE_OK
const sqliteOK = 0

//so:extern SQLITE_ROW
const sqliteRow = 100

//so:extern SQLITE_DONE
const sqliteDone = 101

type sqlite3 struct{}
type sqlite3_stmt struct{}
type sqlite3_value struct{}

type sqlite3_callback func(any, int32, **c.Char, **c.Char) int32
type sqlite3_destructor_type func(any)

func sqlite3_open(filename string, ppDb **sqlite3) int32
func sqlite3_prepare_v2(db *sqlite3, zSql string, nByte int32, ppStmt **sqlite3_stmt, pzTail **c.ConstChar) int32
func sqlite3_step(arg0 *sqlite3_stmt) int32
func sqlite3_finalize(pStmt *sqlite3_stmt) int32
func sqlite3_close(arg0 *sqlite3) int32
func sqlite3_exec(arg0 *sqlite3, sql string, callback sqlite3_callback, arg3 any, errmsg **c.Char) int32

func sqlite3_bind_blob(arg0 *sqlite3_stmt, arg1 int32, arg2 any, n int32, arg4 sqlite3_destructor_type) int32
func sqlite3_bind_double(arg0 *sqlite3_stmt, arg1 int32, arg2 float64) int32
func sqlite3_bind_int(arg0 *sqlite3_stmt, arg1 int32, arg2 int32) int32
func sqlite3_bind_int64(arg0 *sqlite3_stmt, arg1 int32, arg2 int64) int32
func sqlite3_bind_null(arg0 *sqlite3_stmt, arg1 int32) int32
func sqlite3_bind_pointer(arg0 *sqlite3_stmt, arg1 int32, arg2 any, arg3 string, arg4 sqlite3_destructor_type) int32
func sqlite3_bind_text(arg0 *sqlite3_stmt, arg1 int32, arg2 string, arg3 int32, arg4 sqlite3_destructor_type) int32
func sqlite3_bind_value(arg0 *sqlite3_stmt, arg1 int32, arg2 *sqlite3_value) int32
func sqlite3_bind_zeroblob(arg0 *sqlite3_stmt, arg1 int32, n int32) int32

func sqlite3_column_blob(arg0 *sqlite3_stmt, iCol int32) any
func sqlite3_column_bytes(arg0 *sqlite3_stmt, iCol int32) int32
func sqlite3_column_double(arg0 *sqlite3_stmt, iCol int32) float64
func sqlite3_column_int(arg0 *sqlite3_stmt, iCol int32) int32
func sqlite3_column_int64(arg0 *sqlite3_stmt, iCol int32) int64
func sqlite3_column_text(arg0 *sqlite3_stmt, iCol int32) *c.ConstChar
func sqlite3_column_type(arg0 *sqlite3_stmt, iCol int32) int32
func sqlite3_column_value(arg0 *sqlite3_stmt, iCol int32) *sqlite3_value
