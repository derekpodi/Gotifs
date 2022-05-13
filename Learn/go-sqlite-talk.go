//https://www.youtube.com/watch?v=XcAYkriuQ1o
//GopherCon 2021: Ben Johnson - Building Production Applications Using Go & SQLite
package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func run() error {
	db, err := sql.Open("sqlite3", "path/to/db")
	if err != nil {
		return err
	}
	defer db.Close()

	//do database stuff here . . .
}

func TestDB(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	//do testing stuff here . . .
}

//https://github.com/benbjohnson/production-sqlite-go

//SQLite Configs
//Journal mode - configs how sqlite writes transactions		PRAGMA journal_mode = WAL;
//Busy Timeout - set how long write transaction will wait to start PRAGMA busy_timeout = 5000;	(5 seconds)
//Foreign Keys 	PRAGMA foreign)keys = ON;

//SQLite Type System
//5 Types:	INTEGER -integral numbers	REAL-floating-point numbers
//	TEXT -readable text data	BLOB-binary data	NULL-no data
//Strict mode for labeling table column types
//Missing timestamp and decimal type

//Time Values
//ISO8601 / RFC 3339 (2000-01-01T00:00:00Z)
//Unix epoch

//Durability
//1)Regular Backups  $ sqlite3 my.db ".backup 'mybackup.db'"
//aws s3 is cheap to store data, expensive to retrieve. Works well, cheap as backup storage (sqlite as a compressed file)
//2)Litestream  $ litestream replicate db s3://mybucket/db
//continuous streaming backup to s3 -- costs as low as $0.03/month

//Performance
//Most sql databases are B-tree based. Same underlying data structure.
//Different specific optimizations, similar performance profile.
//One key difference is network latency - can be large factor for point queries and simple range queries
//SQLite has no network - embedded
//Horizontal scaling SQLite

//When to avoid SQLite
//if you already have working pipeline, you have long running write transactions, you are running ephemeral serverless environment

//Development experience, no external dependencies, up and running fast, little config needed. Watch out or type system! Timestamps and decimal types can be tricky
//Testing experience - built in support for in memory db, run thousands of tests per second using t.Parallel(). Tailscale sqlite implementation needed for parallelization
//Production experience - regualar backups, stream backups, clustered sqlite using rqlite LiteReplica
//Single node performance is fast, no network latency or serialization overhead. Scales vertically + horizontally
