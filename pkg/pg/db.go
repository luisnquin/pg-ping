package pg

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// DB represents a pingable DB
type DB struct {
	db   *sql.DB
	conf Config
}

// NewDB creates a new DB connection
func NewDB(conf Config) (*DB, error) {
	db, err := sql.Open("postgres", conf.ConnStr())
	if err != nil {
		return nil, err
	}
	return &DB{db: db, conf: conf}, nil
}

// Close a pingable DB
func (db *DB) Close() error {
	return db.db.Close()
}

// PingOnce will execute query only once
func (db *DB) PingOnce() chan SQLResult {
	result := make(chan SQLResult, 1)
	go func() {
		result <- executeQuery(db.db, db.conf.GetQuery())
		close(result)
	}()
	return result
}

// Ping will execute query indefinitely
func (db *DB) Ping() chan SQLResult {
	result := make(chan SQLResult, 10)
	go func() {
		ticker := time.NewTicker(db.conf.GetFrequency())
		for range ticker.C {
			result <- executeQuery(db.db, db.conf.GetQuery())
		}
		close(result)
	}()
	return result
}

// SQLResult tracks sql response and time taken
type SQLResult struct {
	Value          string
	TimeTakenInMS  float64
	Failed         bool
	FailureMessage string
}

func executeQuery(db *sql.DB, query string) SQLResult {
	res := SQLResult{}

	ctx, cancelFunc := context.WithCancel(context.Background())
	tenSecondTimer := time.NewTimer(10 * time.Second)
	go func() {
		<-tenSecondTimer.C
		cancelFunc()
	}()

	start := time.Now()
	rows, err := db.QueryContext(ctx, query)
	res.TimeTakenInMS = time.Since(start).Seconds() * 1000

	if err != nil {
		res.Failed = true
		res.FailureMessage = err.Error()
		return res
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			res.Failed = true
			res.FailureMessage = err.Error()
			return res
		}

		res.Value = name
	}

	return res
}